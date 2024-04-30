package references

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"fmt"
	"regexp"

	"github.com/CycloneDX/cyclonedx-go"
)

type RegexpEnricher struct {
	models.Enricher
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	return true
}

func (e *RegexpEnricher) Enrich(component *cyclonedx.Component) error {
	rules := utils.LoadRules()

	for _, item := range rules {
		r, err := regexp.Compile(item.Rule)

		if err != nil {
			return err
		}

		if r.MatchString(utils.GetRealPurl(component.PackageURL)) {
			for _, ref := range item.References {
				if !hasKey(*component.ExternalReferences, ref.URL, ref.Type) {
					*component.ExternalReferences = append(*component.ExternalReferences, cyclonedx.ExternalReference{
						URL:     ref.URL,
						Type:    cyclonedx.ExternalReferenceType(ref.Type),
						Comment: ref.Comment,
					})
				}
			}
			return nil
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
