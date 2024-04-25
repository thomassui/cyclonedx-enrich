package database

import (
	"cyclonedx-enrich/models"
	"errors"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
			return register()
		}},
	}
}

func connect() (*gorm.DB, error) {
	filename := os.Getenv("DATABASE_FILE")

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		if err := create(filename); err != nil {
			return nil, err
		}
	}

	return gorm.Open(sqlite.Open(filename), &gorm.Config{})
}

func create(filename string) error {
	_, err := os.Create(filename)

	if err != nil {
		return err
	}

	return register()
}
