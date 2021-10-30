package main

import (
	"net/http"

	sentryinit "github.com.buker/go-api-starter/internal/sentryinit"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run()
}

// The engine with all endpoints is now extracted from the main function
func setupServer() *gin.Engine {
	r := gin.Default()
	sentryinit.Init()
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	})

	r.Use(gzip.Gzip(gzip.BestSpeed))
	r.GET("/health", healthEndpoint)

	return r
}

func healthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
