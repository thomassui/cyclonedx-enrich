package utils

import (
	"cyclonedx-enrich/models"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/CycloneDX/cyclonedx-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var isDatabaseInitialized bool
var database *gorm.DB

func ResetDatabase() {
	database = nil
	isDatabaseInitialized = false
}

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
		return nil, err
	}

	return gorm.Open(sqlite.Open(filename), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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

func EnrichDB(component *cyclonedx.Component, preload string, fn func(item *models.Component) error) error {
	db := ConnectDatabase()

	if db == nil {
		return fmt.Errorf("Unable to access database")
	}

	var item *models.Component
	db.Where("purl = ?", GetRealPurl(component.PackageURL)).Preload(preload).First(&item)

	if item != nil {
		return fn(item)
	}

	return nil
}
