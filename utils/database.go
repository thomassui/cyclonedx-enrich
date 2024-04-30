package utils

import (
	"cyclonedx-enrich/models"
	"errors"
	"log/slog"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var isDatabaseInitialized bool
var database *gorm.DB

func ConnectDatabase() *gorm.DB {
	if !isDatabaseInitialized {
		db, err := connect()

		if err != nil {
			log.Warn("unable to connect to database",
				slog.String("error", err.Error()))
		}

		database = db
		isDatabaseInitialized = true
	}

	return database
}

func connect() (*gorm.DB, error) {
	filename := Getenv("DATABASE_FILE", "database.db")

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		if err := create(filename); err != nil {
			return nil, err
		}
	}

	return gorm.Open(sqlite.Open(filename), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func create(filename string) error {
	_, err := os.Create(filename)

	if err != nil {
		return err
	}

	return Register()
}

func Register() error {
	entities := []interface{}{
		&models.License{},
		&models.Component{},
	}
	db, err := connect()

	if err != nil {
		return err
	}

	err = db.AutoMigrate(entities...)

	return err
}
