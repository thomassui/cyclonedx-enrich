package managers

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestMavenEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		{name: "Test with nil package", component: nil, want: true},
		{name: "Test with empty package", component: utils.ComponentEmpty, want: true},
		{name: "Test with cocoapods package", component: utils.ComponentCocoapods, want: true},
		{name: "Test with maven package", component: utils.ComponentMaven, want: false},
		{name: "Test with npm package", component: utils.ComponentNpm, want: true},
		{name: "Test with pypi package", component: utils.ComponentPypi, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &MavenEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("MavenEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMavenEnricher_Enrich(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name    string
		e       *MavenEnricher
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Enrich(tt.args.component); (err != nil) != tt.wantErr {
				t.Errorf("MavenEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
