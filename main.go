package main

import (
	"cyclonedx-enrich/cmd/api"
	"cyclonedx-enrich/cmd/database"
	"cyclonedx-enrich/models"
	"flag"

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
