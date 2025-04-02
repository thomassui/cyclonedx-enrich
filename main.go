package main

import (
	"flag"

	"github.com/fnxpt/cyclonedx-enrich/cmd/api"
	"github.com/fnxpt/cyclonedx-enrich/cmd/database"
	"github.com/fnxpt/cyclonedx-enrich/cmd/sbom"
	"github.com/fnxpt/cyclonedx-enrich/models"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

	parseArguments()
}

func getCommands() []models.Commandable {
	return []models.Commandable{
		database.DatabaseCMD{},
		api.ApiCMD{},
		sbom.SbomCMD{},
	}
}

func parseArguments() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	for _, item := range getCommands() {

		if p, ok := item.(models.Commandable); ok {
			for _, command := range p.Commands() {
				if command.NeedsValue {
					flag.CommandLine.Func(command.Flag, command.Description, command.Handler)
				} else {
					flag.CommandLine.BoolFunc(command.Flag, command.Description, command.Handler)
				}
			}
		}
	}

	flag.Parse()
}
