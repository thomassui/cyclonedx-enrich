package licenses

import (
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestExtractEnricher_Skip(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		e    *ExtractEnricher
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Skip(tt.args.component); got != tt.want {
				t.Errorf("ExtractEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractEnricher_Enrich(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name    string
		e       *ExtractEnricher
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Enrich(tt.args.component); (err != nil) != tt.wantErr {
				t.Errorf("ExtractEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
