package managers

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestCocoapodsEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		{name: "Test with nil package", component: nil, want: true},
		{name: "Test with empty package", component: utils.ComponentEmpty, want: true},
		{name: "Test with cocoapods package", component: utils.ComponentCocoapods, want: false},
		{name: "Test with maven package", component: utils.ComponentMaven, want: true},
		{name: "Test with npm package", component: utils.ComponentMaven, want: true},
		{name: "Test with pypi package", component: utils.ComponentPypi, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &CocoapodsEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("CocoapodsEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCocoapodsEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		{name: "Test with invalid cocoapods package", component: utils.ComponentInvalidCocoapods, wantErr: true},
		{name: "Test with cocoapods package", component: utils.ComponentCocoapods, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &CocoapodsEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("CocoapodsEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_path(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := path(tt.args.name); got != tt.want {
				t.Errorf("path() = %v, want %v", got, tt.want)
			}
		})
	}
}
