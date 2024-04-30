package properties

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

func (e *RegexpEnricher) Category() models.EnricherCategory {
	return "properties"
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	if len(utils.LoadRules()) == 0 {
		return true
	}

	return false
}

func (e *RegexpEnricher) Enrich(component *cyclonedx.Component) error {
	rules := utils.LoadRules()

	for _, item := range rules {
		r, err := regexp.Compile(item.Rule)

		if err != nil {
			return err
		}

		if r.MatchString(utils.GetRealPurl(component.PackageURL)) {
			for key, value := range item.Properties {
				if !hasKey(*component.Properties, key) {
					*component.Properties = append(*component.Properties, cyclonedx.Property{
						Name:  key,
						Value: value,
					})
				}
			}
			return nil
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
