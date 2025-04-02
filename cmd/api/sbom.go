package api

import (
	"log/slog"
	"net/http"

	"github.com/fnxpt/cyclonedx-enrich/cmd/sbom"

	"github.com/gin-gonic/gin"
)

func ParseSBOM(c *gin.Context) {

	request, err := sbom.Enrich(c.Request.Body)

	if err != nil {
		log.Error("Unable to bind sbom",
			slog.String("error", err.Error()))
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, request)
}
