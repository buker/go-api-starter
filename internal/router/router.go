package router

import (
	"github.com/buker/go-api-starter/internal/api"
	"github.com/buker/go-api-starter/internal/middlewares"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ginlog "github.com/onrik/logrus/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/spf13/viper"
)

//Func to create router
func SetupRouter() *gin.Engine {
	if viper.GetString("app.env") == "production" {
		gin.SetMode(gin.ReleaseMode) // Omit this line to enable debug mode
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/metrics"})))
	router.Use(ginlog.Middleware(ginlog.DefaultConfig))
	router.Use(middlewares.CORS())
	// get global Monitor object
	metrics := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	metrics.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	metrics.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	metrics.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	metrics.Use(router)

	// API:v1.0
	tools := router.Group("/tools")
	auth := router.Group("/auth")
	v1 := router.Group("/api/v1/")
	{

		// Tools
		api.Tools(tools.Group(""))
		api.Auth(auth.Group(""))
		api.User(v1.Group("user"))

	}

	return router
}
