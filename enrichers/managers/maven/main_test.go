package maven

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestMavenEnricher_Skip(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		e    *MavenEnricher
		args args
		want bool
	}{
		//TODO: CONTINUE
		{name: "Test with empty component", e: &MavenEnricher{}, args: args{utils.ComponentEmpty}, want: true},
		{name: "Test with component with cocoapods", e: &MavenEnricher{}, args: args{utils.ComponentCocoapods}, want: true},
		{name: "Test with component with maven", e: &MavenEnricher{}, args: args{utils.ComponentMaven}, want: false},
		{name: "Test with component with npm", e: &MavenEnricher{}, args: args{utils.ComponentNpm}, want: true},
		{name: "Test with component with pypi", e: &MavenEnricher{}, args: args{utils.ComponentPypi}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Skip(tt.args.component); got != tt.want {
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
		//TODO: CONTINUE
		{name: "Test with component with maven", e: &MavenEnricher{}, args: args{utils.ComponentMaven}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Enrich(tt.args.component); (err != nil) != tt.wantErr {
				t.Errorf("MavenEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
