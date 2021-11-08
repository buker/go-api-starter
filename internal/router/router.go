package router

import (
	"github.com/gin-contrib/gzip"

	"github.com/buker/go-api-starter/cmd/api"
	"github.com/buker/go-api-starter/internal/config"
	"github.com/buker/go-api-starter/internal/controller"
	"github.com/buker/go-api-starter/internal/middleware"
	"github.com/gin-gonic/gin"
	ginlog "github.com/onrik/logrus/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

var configure = config.Config()

//Func to create router
func SetupRouter() *gin.Engine {
	if configure.Server.ServerEnv == "production" {
		gin.SetMode(gin.ReleaseMode) // Omit this line to enable debug mode
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/metrics"})))
	router.Use(ginlog.Middleware(ginlog.DefaultConfig))
	router.Use(middleware.CORS())
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
	v1 := router.Group("/api/v1/")
	{
		// Register - no JWT required
		v1.POST("register", controller.CreateUserAuth)

		// Login - app issues JWT
		v1.POST("login", controller.Login)

		// User
		api.Users(v1.Group("users"))

		// Tools
		api.Tools(v1.Group("tools"))
	}

	return router
}
