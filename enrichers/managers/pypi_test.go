package managers

import (
	"testing"

	"github.com/fnxpt/cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestPyPiEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		{name: "Test with nil package", component: nil, want: true},
		{name: "Test with empty package", component: utils.ComponentEmpty, want: true},
		{name: "Test with cocoapods package", component: utils.ComponentCocoapods, want: true},
		{name: "Test with maven package", component: utils.ComponentMaven, want: true},
		{name: "Test with npm package", component: utils.ComponentNpm, want: true},
		{name: "Test with pypi package", component: utils.ComponentPypi, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &PyPiEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("PyPiEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPyPiEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		{name: "Test with pypi package", component: utils.ComponentPypi, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &PyPiEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("PyPiEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
