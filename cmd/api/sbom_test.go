package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParseSBOM(t *testing.T) {
	tests := []struct {
		name string
		c    *gin.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSBOM(tt.c)
		})
	}
}
