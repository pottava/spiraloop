package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/pottava/spiraloop/api/generated/restapi"
	"github.com/pottava/spiraloop/api/generated/restapi/operations"
	"github.com/pottava/spiraloop/api/logs"
)

func main() {
	logs.Debug("start api", nil, &logs.Map{})

	// ----------------------------------------------------------------------------------------
	//  Copied from generated/cmd/spiraloop-server/main.go
	// ----------------------------------------------------------------------------------------

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSpiraloopAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "spiraloop"
	parser.LongDescription = "For better release observability."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err) // nolint:gocritic
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
	// ----------------------------------------------------------------------------------------
}
