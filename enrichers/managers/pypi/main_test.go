package pypi

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestPyPiEnricher_Skip(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		e    *PyPiEnricher
		args args
		want bool
	}{
		//TODO: CONTINUE
		{name: "Test with empty component", e: &PyPiEnricher{}, args: args{utils.ComponentEmpty}, want: true},
		{name: "Test with component with cocoapods", e: &PyPiEnricher{}, args: args{utils.ComponentCocoapods}, want: true},
		{name: "Test with component with maven", e: &PyPiEnricher{}, args: args{utils.ComponentMaven}, want: true},
		{name: "Test with component with npm", e: &PyPiEnricher{}, args: args{utils.ComponentNpm}, want: true},
		{name: "Test with component with pypi", e: &PyPiEnricher{}, args: args{utils.ComponentPypi}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Skip(tt.args.component); got != tt.want {
				t.Errorf("PyPiEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPyPiEnricher_Enrich(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name    string
		e       *PyPiEnricher
		args    args
		wantErr bool
	}{
		//TODO: CONTINUE
		{name: "Test with component with pypi", e: &PyPiEnricher{}, args: args{utils.ComponentPypi}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Enrich(tt.args.component); (err != nil) != tt.wantErr {
				t.Errorf("PyPiEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
