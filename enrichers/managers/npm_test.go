package managers

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestNPMEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		{name: "Test with nil package", component: nil, want: true},
		{name: "Test with empty package", component: utils.ComponentEmpty, want: true},
		{name: "Test with cocoapods package", component: utils.ComponentCocoapods, want: true},
		{name: "Test with maven package", component: utils.ComponentMaven, want: true},
		{name: "Test with npm package", component: utils.ComponentMaven, want: false},
		{name: "Test with pypi package", component: utils.ComponentPypi, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NPMEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("NPMEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNPMEnricher_Enrich(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		//TODO
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NPMEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("NPMEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
