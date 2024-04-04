package enrich

import (
	"regexp"

	"github.com/CycloneDX/cyclonedx-go"
)

func Enrich(sbom *cyclonedx.BOM, patterns map[string][]string, force bool) {
	changed := *sbom.Components

	for pattern, license := range patterns {
		r, _ := regexp.Compile(pattern)

		for idx, comp := range changed {
			if r.MatchString(comp.PackageURL) {

				if comp.Licenses == nil || force {
					changed[idx].Licenses = getLicensesObject(license)
				} else {
					value := []cyclonedx.LicenseChoice(*comp.Licenses)
					if len(value) == 0 || force {
						changed[idx].Licenses = getLicensesObject(license)
					}
				}
			}
		}
	}

	sbom.Components = &changed
}

func getLicensesObject(input []string) *cyclonedx.Licenses {
	licenses := make(cyclonedx.Licenses, 0)

	for _, id := range input {
		licenses = append(licenses, cyclonedx.LicenseChoice{
			License: &cyclonedx.License{Name: id},
		})
	}
	return &licenses
}
