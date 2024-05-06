package managers

import (
	"cyclonedx-enrich/utils"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

func skip(component *cyclonedx.Component, prefix string) bool {
	if component == nil || !strings.HasPrefix(utils.GetRealPurl(component.PackageURL), prefix) {
		return true
	}
	if component.Licenses != nil {
		//SKIP
		return true
	}
	return false
}

func enrich(component *cyclonedx.Component, licenses []string, hashes map[string]string, references map[string]string, properties map[string]string) {
	utils.SetLicense(component, licenses)
}

func request[T any](url string, fn func(io.ReadCloser) (*T, error)) (*T, error) {
	r, err := utils.Request(url)
	if err != nil {
		log.Error("error with request",
			slog.String("url", url),
			slog.String("error", err.Error()))
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code %d", r.StatusCode)
	}

	return fn(r.Body)

}

func parseJSON[T any](input io.ReadCloser) (*T, error) {
	item := new(T)
	err := json.NewDecoder(input).Decode(item)

	if err != nil {
		log.Error("cannot unmarshal",
			slog.String("error", err.Error()))
		return nil, err
	}

	return item, nil
}
