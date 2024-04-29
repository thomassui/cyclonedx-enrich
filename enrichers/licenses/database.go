package licenses

import (
	"cyclonedx-enrich/models"
	"fmt"
	"log/slog"
	"os"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseEnricher struct {
	models.Enricher
	isInitialized bool
	database      *gorm.DB
}

func connect() (*gorm.DB, error) {
	filename := os.Getenv("DATABASE_FILE")
	return gorm.Open(sqlite.Open(filename), &gorm.Config{})
}

func (e *DatabaseEnricher) Category() models.EnricherCategory {
	return "license"
}

func (e *DatabaseEnricher) Skip(component *cyclonedx.Component) bool {
	// TODO: VALIDATE IF IT HAS EXPRESSION OR LICENSE OBJECT
	if len(component.PackageURL) == 0 || component.Licenses != nil {
		//SKIP
		return true
	}
	return false
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	if !e.isInitialized {
		db, err := connect()

		if err != nil {
			log.Error("unable to load rules",
				slog.String("error", err.Error()))
		}

		e.database = db
		e.isInitialized = true
	}

	if e.database == nil {
		return nil
	}

	var item *models.Component
	e.database.Where("purl = ?", utils.GetRealPurl(component.PackageURL)).Preload("Licenses").First(&item)

	if item != nil {
		licenses := make([]string, 0)

		for _, license := range item.Licenses {
			licenses = append(licenses, license.License)
		}

		if len(licenses) > 0 {
			utils.SetLicense(component, licenses)
			return nil
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
