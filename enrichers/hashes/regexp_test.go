package hashes

import (
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/fnxpt/cyclonedx-enrich/utils"
)

func TestRegexpEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		validEnv  bool
		component cyclonedx.Component
		want      bool
	}{
		{name: "Test without rules with empty component", validEnv: false, component: *utils.ComponentEmpty, want: true},
		{name: "Test without rules with component with data", validEnv: false, component: *utils.ComponentWithData, want: true},
		{name: "Test without rules with component without data", validEnv: false, component: *utils.ComponentWithoutData, want: true},
		{name: "Test with empty component", validEnv: true, component: *utils.ComponentEmpty, want: true},
		{name: "Test with component with data", validEnv: true, component: *utils.ComponentWithData, want: true},
		{name: "Test with component without data", validEnv: true, component: *utils.ComponentWithoutData, want: false},
	}
	for _, tt := range tests {
		teardown := setup(t, tt.validEnv)
		defer teardown(t)
		t.Run(tt.name, func(t *testing.T) {
			e := &RegexpEnricher{}
			if got := e.Skip(&tt.component); got != tt.want {
				t.Errorf("RegexpEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexpEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		validEnv  bool
		component cyclonedx.Component
		wantErr   bool
	}{
		{name: "Test without rules with component without data", validEnv: false, component: *utils.ComponentWithoutData, wantErr: true},
		{name: "Test with component without data", validEnv: true, component: *utils.ComponentWithoutData, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t, tt.validEnv)
			defer teardown(t)
			e := &RegexpEnricher{}
			if err := e.Enrich(&tt.component); (err != nil) != tt.wantErr {
				t.Errorf("RegexpEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
