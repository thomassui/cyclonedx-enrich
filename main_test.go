package main

import (
	"reflect"
	"testing"

	"github.com/fnxpt/cyclonedx-enrich/models"
	_ "github.com/joho/godotenv/autoload"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_getCommands(t *testing.T) {
	tests := []struct {
		name string
		want []models.Commandable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseArguments(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseArguments()
		})
	}
}
