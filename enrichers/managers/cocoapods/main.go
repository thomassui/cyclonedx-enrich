package cocoapods

import (
	"crypto/md5"
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

var endpoint = "https://cdn.cocoapods.org"
var client = &http.Client{Timeout: 10 * time.Second}
var pathLenght = 3 //TODO: GET THIS FROM https://cdn.cocoapods.org/CocoaPods-version.yml"

type CocoapodsEnricher struct {
	models.Enricher
}

func (e *CocoapodsEnricher) Category() models.EnricherCategory {
	return "managers"
}

func (e *CocoapodsEnricher) Skip(component *cyclonedx.Component) bool {
	if !strings.HasPrefix(utils.GetRealPurl(component.PackageURL), "pkg:cocoapods/") {
		return true
	}
	if component.Licenses != nil {
		//SKIP
		return true
	}
	return false
}

func (e *CocoapodsEnricher) Enrich(component *cyclonedx.Component) error {
	url := fmt.Sprintf("%s/Specs/%s/%s/%s/%s.podspec.json", endpoint, path(component.Name), component.Name, component.Version, component.Name)

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

func path(name string) string {
	hash := md5.Sum([]byte(name))
	hashString := hex.EncodeToString(hash[:])
	value := make([]string, pathLenght)

	for idx := 0; idx < pathLenght; idx++ {
		value[idx] = string(hashString[idx])
	}

	return strings.Join(value, "/")
}
