package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-openapi/swag"
	commands "github.com/pottava/spiraloop/cli"
	"github.com/pottava/spiraloop/cli/conf"
	"github.com/pottava/spiraloop/cli/internal/log"
	cli "gopkg.in/alecthomas/kingpin.v2"
)

// for compile flags
var (
	ver    = "dev"
	commit string
	date   string
)

func main() {

	app := cli.New("spiraloop", "A tool for release observability")
	if len(commit) > 0 && len(date) > 0 {
		app.Version(fmt.Sprintf("%s-%s (built at %s)", ver, commit, date))
	} else {
		app.Version(ver)
	}

	common := &conf.CommonConfig{}
	common.APIEndpoint = app.Flag("api-endpoint", "An endpoint of spiraloop apis.").
		Envar("API_ENDPOINT").Default("localhost:8080").String()
	common.APIKey = app.Flag("key", "An endpoint of spiraloop apis.").
		Envar("API_KEY").Required().String()
	common.AppVersion = ver
	if len(commit) > 0 && len(date) > 0 {
		common.AppVersion = fmt.Sprintf("%s-%s (built at %s)", ver, commit, date)
	}
	common.ExtendedOutput = app.Flag("extended-output", "If it's True, meta data returns as well.").
		Envar("EXTENDED_OUTPUT").Default("false").Bool()
	common.IsDebugMode = os.Getenv("APP_DEBUG") == "1"

	// commands
	startconf := &conf.StartConfig{}
	startconf.Common = common
	start := app.Command("start", "Notify a process has started.")
	startconf.Platform = start.Flag("platform", "Platform the process is running on.").
		Envar("PLATFORM").Default("git").String()
	startconf.Asynchronous = start.Flag("async", "If it's True, the app does not wait for the request.").
		Envar("ASYNC").Default("false").Bool()

	successconf := &conf.SuccessConfig{}
	successconf.Common = common
	success := app.Command("success", "Notify the process has done successfully.")

	// Cancel
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
		os.Exit(1)
	}()

	switch cli.MustParse(app.Parse(os.Args[1:])) {
	case start.FullCommand():

		// Execute
		out, err := commands.Start(ctx, startconf)
		if err != nil {
			log.Errors.Fatal(err)
			return
		}
		if swag.BoolValue(startconf.Asynchronous) {
			if swag.BoolValue(common.ExtendedOutput) {
				log.PrintJSON(struct {
					RequestID string              `json:"RequestID"`
					Meta      commands.OutputMeta `json:"meta"`
				}{
					RequestID: out.RequestID,
					Meta:      out.Meta,
				})
			} else {
				log.PrintJSON(struct {
					RequestID string `json:"RequestID"`
				}{
					RequestID: out.RequestID,
				})
			}
		} else {
			if swag.BoolValue(common.ExtendedOutput) {
				out.SyncLogs["meta"] = out.Meta
			}
			log.PrintJSON(out.SyncLogs)
		}
		os.Exit(int(swag.Int64Value(out.ExitCode)))

	case success.FullCommand():

		// Execute
		out, err := commands.Success(ctx, successconf)
		if err != nil {
			log.Errors.Fatal(err)
			return
		}
		if swag.BoolValue(common.ExtendedOutput) {
			out.SyncLogs["meta"] = out.Meta
		}
		log.PrintJSON(out.SyncLogs)
		os.Exit(int(swag.Int64Value(out.ExitCode)))
	}
}
