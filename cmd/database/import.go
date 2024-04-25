package database

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/CycloneDX/cyclonedx-go"
)

func importFile(filename string) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	return importComponent(data)
}

func importComponent(data []byte) error {

	var bom *cyclonedx.BOM

	err := json.Unmarshal(data, &bom)

	if err != nil {
		return err
	}

	if bom.Metadata == nil || bom.Metadata.Component == nil {
		return fmt.Errorf("BOM doesn't contain a main component")
	}

	item := models.Component{
		Purl:       utils.GetRealPurl(bom.Metadata.Component.PackageURL),
		Licenses:   getLicenses(bom.Metadata.Component),
		Properties: getProperties(bom.Metadata.Component),
		Hashes:     getHashes(bom.Metadata.Component),
		References: getReferences(bom.Metadata.Component),
	}

	db, err := connect()
	if err != nil {
		return err
	}

	tx := db.Save(item)

	return tx.Error
}

func getLicenses(component *cyclonedx.Component) []models.License {
	output := make([]models.License, 0)

	if component.Licenses == nil {
		return output
	}

	for _, item := range *component.Licenses {
		//TODO: EXPRESSION
		if item.License != nil {
			if len(item.License.ID) > 0 {
				output = append(output, models.License{License: item.License.ID})
			} else if len(item.License.Name) > 0 {
				output = append(output, models.License{License: item.License.Name})
			}
		}
	}

	return output
}

func getProperties(component *cyclonedx.Component) []models.Property {
	output := make([]models.Property, 0)

	if component.Properties == nil {
		return output
	}

	for _, item := range *component.Properties {
		output = append(output, models.Property{Name: item.Name, Value: item.Value})
	}

	return output
}

func getReferences(component *cyclonedx.Component) []models.Reference {
	output := make([]models.Reference, 0)

	if component.ExternalReferences == nil {
		return output
	}

	for _, item := range *component.ExternalReferences {
		output = append(output, models.Reference{URL: item.URL, Type: string(item.Type), Comment: item.Comment})
	}

	return output
}

func getHashes(component *cyclonedx.Component) []models.Hash {
	output := make([]models.Hash, 0)

	if component.Hashes == nil {
		return output
	}

	for _, item := range *component.Hashes {
		output = append(output, models.Hash{Name: string(item.Algorithm), Value: item.Value})
	}

	return output
}
