package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"github.com/quintsys/ga-exporter/api/restapi"
	"github.com/quintsys/ga-exporter/api/restapi/operations"
)

func main() {
	// Load the embedded swagger spec
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Create the api and server
	api := operations.NewGaexporterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		_ = server.Shutdown()
	}()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = swaggerSpec.Spec().Info.Title
	parser.LongDescription = "Exports visitor's information from Google Analytics"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
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

	// Start the server
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
