package database

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
)

type DatabaseCMD struct {
	models.Commandable
}

func (c DatabaseCMD) Commands() []models.Command {

	return []models.Command{
		{Flag: "database-import", Description: "Imports cyclonedx component into database", NeedsValue: true, Handler: func(value string) error {
			return importFile(value)
		}},
		{Flag: "database-download", Description: "Downloads database from source", Handler: func(value string) error {
			return download()
		}},
		{Flag: "database-register", Description: "Registers database entities", Handler: func(value string) error {
			return utils.Register()
		}},
	}
}
