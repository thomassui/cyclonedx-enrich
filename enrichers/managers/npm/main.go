package npm

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

var endpoint = "https://registry.npmjs.org"
var client = &http.Client{Timeout: 10 * time.Second}

type NPMEnricher struct {
	models.Enricher
}

func (e *NPMEnricher) Category() models.EnricherCategory {
	return "managers"
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

	r, err := client.Get(url)
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
