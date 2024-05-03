package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthorizeRequest(t *testing.T) {
	tests := []struct {
		name     string
		isPublic bool
		want     gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthorizeRequest(tt.isPublic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandToken(t *testing.T) {
	tests := []struct {
		name    string
		l       int
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandToken(tt.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RandToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Setup(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Setup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRoutes(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 []Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getRoutes()
			if got != tt.want {
				t.Errorf("getRoutes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getRoutes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClearCache(t *testing.T) {
	tests := []struct {
		name string
		c    *gin.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearCache(tt.c)
		})
	}
}
