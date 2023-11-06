package enrich

import (
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
)

func Enrich(sbom *cyclonedx.BOM, pattern string, license string) {
	changed := *sbom.Components
	for idx, comp := range changed {
		if strings.Contains(comp.PackageURL, pattern) {

			if comp.Licenses == nil {
				changed[idx].Licenses = &cyclonedx.Licenses{{License: &cyclonedx.License{ID: license}}}
			} else {
				value := []cyclonedx.LicenseChoice(*comp.Licenses)
				if len(value) == 0 {
					changed[idx].Licenses = &cyclonedx.Licenses{{License: &cyclonedx.License{ID: license}}}
				}
			}
		}
	}

	sbom.Components = &changed
}
