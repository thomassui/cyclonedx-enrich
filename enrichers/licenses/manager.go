package licenses

import (
	"fmt"
	"log/slog"
	"strings"

	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

type ManagerEnricher struct {
	models.Enricher
}

func (e *ManagerEnricher) Category() models.EnricherCategory {
	return "license"
}

func (e *ManagerEnricher) Skip(component *cyclonedx.Component) bool {
	// TODO: VALIDATE IF IT HAS EXPRESSION OR LICENSE OBJECT
	if len(component.PackageURL) == 0 || component.Licenses != nil {
		//SKIP
		return true
	}
	return false
}

func (e *ManagerEnricher) Enrich(component *cyclonedx.Component) error {
	// GET FROM PACKAGE MANAGER
	purl := strings.TrimPrefix(utils.GetRealPurl(component.PackageURL), "pkg:maven/")

	log.Debug("parsing component",
		slog.String("purl", purl))

	// searchArtifact(purl, 1)
	// panic("here")

	return fmt.Errorf("component doesn't met criteria")
}
