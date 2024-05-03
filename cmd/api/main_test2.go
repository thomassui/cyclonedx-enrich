package api

import (
	"cyclonedx-enrich/models"
	"reflect"
	"testing"
)

func TestApiCMD_Commands(t *testing.T) {
	tests := []struct {
		name string
		c    ApiCMD
		want []models.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Commands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiCMD.Commands() = %v, want %v", got, tt.want)
			}
		})
	}
}
