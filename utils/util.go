package utils

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/gregjones/httpcache"
)

var tp = httpcache.NewMemoryCacheTransport()
var client = &http.Client{Timeout: 10 * time.Second, Transport: tp}

func GetRealPurl(purl string) string {
	u, _ := url.Parse(purl)

	return Decoded(strings.TrimSuffix(strings.TrimSuffix(purl, "#"+u.Fragment), "?"+u.RawQuery))
}

func Decoded(value string) string {
	decodedValue, err := url.QueryUnescape(value)

	if err != nil {
		return value
	}

	return decodedValue
}

func SetLicense(component *cyclonedx.Component, licenseNames []string) {
	licenses := make([]cyclonedx.LicenseChoice, 0)

	for _, item := range licenseNames {
		licenses = append(licenses, cyclonedx.LicenseChoice{License: &cyclonedx.License{Name: item}})
	}

	component.Licenses = (*cyclonedx.Licenses)(&licenses)
}

func Request(url string) (resp *http.Response, err error) {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cache-Control", "stale-if-error")
	r, err := tp.RoundTrip(req)

	if r.StatusCode == 301 {
		newURL, _ := r.Location()
		return Request(newURL.String())
	}

	return r, err
}

func ReadFile(filename string, fn func(*os.File) error) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	return fn(file)
}

func Getenv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}
