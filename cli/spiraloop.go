package etr

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/google/uuid"
	config "github.com/pottava/spiraloop/cli/conf"
	"github.com/pottava/spiraloop/cli/internal/log"
)

var (
	RequestID string
)

func init() {
	RequestID = fmt.Sprintf("spiraloop-%s", uuid.New().String())
}

// Start runs the docker image on Amazon ECS
func Start(ctx context.Context, conf *config.StartConfig) (output *Output, err error) {
	if conf.Common.IsDebugMode {
		log.PrintJSON(conf)
	}
	defer func() {
		if err := recover(); err != nil {
			if os.Getenv("APP_DEBUG") == "1" {
				debug.PrintStack()
			}
			log.Errors.Fatal(err)
		}
	}()
	logs := map[string]interface{}{}
	logs["status"] = "ok"
	return &Output{SyncLogs: logs}, nil
}

// Success stops the Fargate container on Amazon ECS
func Success(ctx context.Context, conf *config.SuccessConfig) (output *Output, err error) {
	defer func() {
		if err := recover(); err != nil {
			if os.Getenv("APP_DEBUG") == "1" {
				debug.PrintStack()
			}
			log.Errors.Fatal(err)
		}
	}()
	logs := map[string]interface{}{}
	logs["status"] = "ok"
	return &Output{SyncLogs: logs}, nil
}
