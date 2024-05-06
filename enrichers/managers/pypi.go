package managers

import (
	"cyclonedx-enrich/models"
	"fmt"

	"github.com/CycloneDX/cyclonedx-go"
)

var pypiEndpoint = "https://pypi.org/pypi"

type PyPiEnricher struct {
	models.Enricher
}

func (e *PyPiEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component, "pkg:pypi/")
}

func (e *PyPiEnricher) Enrich(component *cyclonedx.Component) error {
	url := fmt.Sprintf("%s/%s/%s/json", pypiEndpoint, component.Name, component.Version)

	item, err := request[PyPiPackage](url, parseJSON)

	if err != nil {
		return err
	}

	if item != nil {
		licenses := make([]string, 0)
		hashes := make(map[string]string)
		references := make(map[string]string)
		properties := make(map[string]string)

		if item.Info.License != nil {
			licenses = append(licenses, *item.Info.License)
		}

		enrich(component, licenses, hashes, references, properties)
	}

	//TODO: USE MORE DATA

	return fmt.Errorf("component doesn't met criteria")
}
