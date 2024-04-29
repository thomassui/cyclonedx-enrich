package pypi

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

var endpoint = "https://pypi.org/pypi"
var client = &http.Client{Timeout: 10 * time.Second}

type PyPiEnricher struct {
	models.Enricher
}

func (e *PyPiEnricher) Category() models.EnricherCategory {
	return "managers"
}

func (e *PyPiEnricher) Skip(component *cyclonedx.Component) bool {
	if !strings.HasPrefix(utils.GetRealPurl(component.PackageURL), "pkg:pypi/") {
		return true
	}
	if component.Licenses != nil {
		//SKIP
		return true
	}
	return false
}

func (e *PyPiEnricher) Enrich(component *cyclonedx.Component) error {
	url := fmt.Sprintf("%s/%s/%s/json", endpoint, component.Name, component.Version)

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

	if item != nil {
		if item.Info.License != nil {
			utils.SetLicense(component, []string{*item.Info.License})
		}
	}

	//TODO: USE MORE DATA

	return fmt.Errorf("component doesn't met criteria")
}
