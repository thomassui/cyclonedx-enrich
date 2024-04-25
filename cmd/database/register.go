package database

import "cyclonedx-enrich/models"

func register() error {
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
