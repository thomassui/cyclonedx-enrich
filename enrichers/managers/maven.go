package managers

import (
	"cyclonedx-enrich/models"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/vifraa/gopom"
)

var mavenEndpoint = "https://search.maven.org/remotecontent?filepath="

type MavenEnricher struct {
	models.Enricher
}

func (e *MavenEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component, "pkg:maven/")
}

func (e *MavenEnricher) Enrich(component *cyclonedx.Component) error {
	url := fmt.Sprintf("%s%s/%s/%s/%s-%s.pom", mavenEndpoint, strings.ReplaceAll(component.Group, ".", "/"), component.Name, component.Version, component.Name, component.Version)

	item, err := request(url, parsePom)

	if err != nil {
		return err
	}

	if item != nil {
		licenses := make([]string, 0)
		hashes := make(map[string]string)
		references := make(map[string]string)
		properties := make(map[string]string)

		if item.Licenses != nil {
			for _, item := range *item.Licenses {
				licenses = append(licenses, *item.Name)
			}
		}

		enrich(component, licenses, hashes, references, properties)
	}

	//TODO: USE MORE DATA

	return fmt.Errorf("component doesn't met criteria")
}

func parsePom(input io.ReadCloser) (*gopom.Project, error) {
	item, err := gopom.ParseFromReader(input)

	if err != nil {
		log.Error("cannot unmarshal",
			slog.String("error", err.Error()))
		return nil, err
	}

	return item, nil
}
