package properties

import (
	"testing"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestRegexpEnricher_Skip(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		e    *RegexpEnricher
		args args
		want bool
	}{
		//TODO: CONTINUE
		{name: "Test with empty component", e: &RegexpEnricher{}, args: args{utils.ComponentEmpty}, want: true},
		{name: "Test with component with data", e: &RegexpEnricher{}, args: args{utils.ComponentWithData}, want: true},
		{name: "Test with component without data", e: &RegexpEnricher{}, args: args{utils.ComponentWithoutData}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Skip(tt.args.component); got != tt.want {
				t.Errorf("RegexpEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexpEnricher_Enrich(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name    string
		e       *RegexpEnricher
		args    args
		wantErr bool
	}{
		//TODO: VALIDATE IF DATA WAS ADDED
		{name: "Test with component without data", e: &RegexpEnricher{}, args: args{utils.ComponentWithoutData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Enrich(tt.args.component); (err != nil) != tt.wantErr {
				t.Errorf("RegexpEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
