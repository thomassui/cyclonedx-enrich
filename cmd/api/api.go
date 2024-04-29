package api

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method   string
	Path     string
	IsPublic bool
	Expires  time.Duration
	Handler  func(c *gin.Context)
}

var token = os.Getenv("APP_TOKEN")
var log = slog.Default()
var storeCache persistence.CacheStore

func AuthorizeRequest(isPublic bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: DEFER???
		if c.Request.Method == "OPTIONS" {
			return
		}

		if len(token) == 0 {
			return
		}

		authHeaderValue := c.GetHeader("X-Api-Key")

		if token != authHeaderValue {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

// RandToken generates a random @l length token.
func RandToken(l int) (string, error) {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func Setup() *gin.Engine {

	router := gin.Default()

	token, err := RandToken(64)
	if err != nil {
		log.Error("unable to generate random token: ",
			slog.String("error", err.Error()))
		panic("error")
	}
	store := cookie.NewStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})

	router.Use(sessions.Sessions("goquestsession", store))
	log.Debug("Initializing routes")

	if storeCache != nil {
		router.Handle(http.MethodDelete, "/admin/cache", AuthorizeRequest(false), ClearCache)
	}

	groupName, routes := getRoutes()

	group := router.Group(groupName)
	// 	group.Use(p.Options())

	for _, route := range routes {
		handlers := []gin.HandlerFunc{}

		handlers = append(handlers, AuthorizeRequest(route.IsPublic))

		if p.storeCache != nil && route.Expires > 0 {
			handlers = append(handlers, cache.CachePageAtomic(storeCache, route.Expires, route.Handler))
		} else {
			handlers = append(handlers, route.Handler)
		}

		group.Handle(route.Method, route.Path, handlers...)
	}
	// }

	return router
}

func getRoutes() (string, []Route) {
	return "/sbom", []Route{
		{Method: http.MethodPost, Path: "/enrich", Handler: ParseSBOM},
	}
}

func ClearCache(c *gin.Context) {
	err := storeCache.Flush()
	if err != nil {
		log.Error("Unable to flush cache",
			slog.String("error", err.Error()))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusAccepted)
}
