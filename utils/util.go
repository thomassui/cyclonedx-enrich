package utils

import (
	"net/url"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
)

func GetRealPurl(purl string) string {
	u, _ := url.Parse(purl)

	return strings.TrimSuffix(purl, "?"+u.RawQuery)

}

func SetLicense(component *cyclonedx.Component, licenseNames []string) {
	licenses := make([]cyclonedx.LicenseChoice, 0)

	for _, item := range licenseNames {
		licenses = append(licenses, cyclonedx.LicenseChoice{License: &cyclonedx.License{Name: item}})
	}

	component.Licenses = (*cyclonedx.Licenses)(&licenses)
}
