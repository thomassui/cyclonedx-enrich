package maven

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"fmt"
	"log/slog"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/vifraa/gopom"
)

var log = slog.Default()

var endpoint = "https://search.maven.org/remotecontent?filepath="

type MavenEnricher struct {
	models.Enricher
}

func (e *MavenEnricher) Category() models.EnricherCategory {
	return "managers"
}

func (e *MavenEnricher) Skip(component *cyclonedx.Component) bool {
	if !strings.HasPrefix(utils.GetRealPurl(component.PackageURL), "pkg:maven/") {
		return true
	}
	if component.Licenses != nil {
		//SKIP
		return true
	}

	return false
}

func (e *MavenEnricher) Enrich(component *cyclonedx.Component) error {
	url := fmt.Sprintf("%s%s/%s/%s/%s-%s.pom", endpoint, strings.ReplaceAll(component.Group, ".", "/"), component.Name, component.Version, component.Name, component.Version)

	r, err := utils.Request(url)

	if err != nil {
		log.Error("error with request",
			slog.String("package", component.PackageURL),
			slog.String("url", url),
			slog.String("error", err.Error()))
		return err
	}
	defer r.Body.Close()

	parsedPom, err := gopom.ParseFromReader(r.Body)

	if err != nil {
		log.Error("cannot unmarshal",
			slog.String("package", component.PackageURL),
			slog.String("url", url),
			slog.String("error", err.Error()))
		return err
	}

	if parsedPom != nil {

		if parsedPom.Licenses != nil {
			licenses := make([]string, 0)

			for _, item := range *parsedPom.Licenses {
				licenses = append(licenses, *item.Name)
			}

			utils.SetLicense(component, licenses)
		}

	}

	//TODO: USE MORE DATA

	return fmt.Errorf("component doesn't met criteria")
}
