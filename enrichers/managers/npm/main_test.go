package npm

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestNPMEnricher_Skip(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		e    *NPMEnricher
		args args
		want bool
	}{
		//TODO: CONTINUE
		{name: "Test with empty component", e: &NPMEnricher{}, args: args{utils.ComponentEmpty}, want: true},
		{name: "Test with component with cocoapods", e: &NPMEnricher{}, args: args{utils.ComponentCocoapods}, want: true},
		{name: "Test with component with maven", e: &NPMEnricher{}, args: args{utils.ComponentMaven}, want: true},
		{name: "Test with component with npm", e: &NPMEnricher{}, args: args{utils.ComponentNpm}, want: false},
		{name: "Test with component with pypi", e: &NPMEnricher{}, args: args{utils.ComponentPypi}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Skip(tt.args.component); got != tt.want {
				t.Errorf("NPMEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNPMEnricher_Enrich(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name    string
		e       *NPMEnricher
		args    args
		wantErr bool
	}{
		//TODO: CONTINUE
		{name: "Test with component with npm", e: &NPMEnricher{}, args: args{utils.ComponentNpm}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Enrich(tt.args.component); (err != nil) != tt.wantErr {
				t.Errorf("NPMEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
