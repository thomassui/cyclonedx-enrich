package npm

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

var endpoint = "https://registry.npmjs.org"

type NPMEnricher struct {
	models.Enricher
}

func (e *NPMEnricher) Skip(component *cyclonedx.Component) bool {
	if !strings.HasPrefix(utils.GetRealPurl(component.PackageURL), "pkg:npm/") {
		return true
	}
	if component.Licenses != nil {
		//SKIP
		return true
	}
	return false
}

func (e *NPMEnricher) Enrich(component *cyclonedx.Component) error {
	var url string

	if len(component.Group) > 0 {
		url = fmt.Sprintf("%s/%s/%s/%s", endpoint, utils.Decoded(component.Group), component.Name, component.Version)
	} else {
		url = fmt.Sprintf("%s/%s/%s", endpoint, component.Name, component.Version)
	}

	r, err := utils.Request(url)
	if err != nil {
		log.Error("error with request",
			slog.String("package", component.PackageURL),
			slog.String("url", url),
			slog.String("error", err.Error()))
		return err
	}
	defer r.Body.Close()

	item := &Package{}
	err = json.NewDecoder(r.Body).Decode(item)

	if err != nil {
		fmt.Println(url)
		log.Error("cannot unmarshal",
			slog.String("package", component.PackageURL),
			slog.String("url", url),
			slog.String("error", err.Error()))
		return err
	}

	if item.License != nil {
		utils.SetLicense(component, []string{*item.License})
	}

	//TODO: USE MORE DATA

	return fmt.Errorf("component doesn't met criteria")
}
