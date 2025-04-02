package api

import "github.com/fnxpt/cyclonedx-enrich/models"

type ApiCMD struct {
	models.Commandable
}

func (c ApiCMD) Commands() []models.Command {

	return []models.Command{
		{Flag: "server", Description: "Starts server", Handler: func(value string) error {
			router := Setup()

			err := router.Run()
			log.Info("Running server...")
			return err
		}},
	}
}
