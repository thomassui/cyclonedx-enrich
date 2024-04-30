package licenses

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
	return "license"
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	if len(utils.LoadRules()) == 0 {
		return true
	}

	// TODO: VALIDATE IF IT HAS EXPRESSION OR LICENSE OBJECT
	if len(component.PackageURL) == 0 || component.Licenses != nil {
		//SKIP
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
			utils.SetLicense(component, item.Licenses)
			return nil
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
