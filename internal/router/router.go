package router

import (
	"github.com/buker/go-api-starter/docs"
	"github.com/buker/go-api-starter/internal/api"
	"github.com/buker/go-api-starter/internal/middlewares"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ginlog "github.com/onrik/logrus/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	docs.SwaggerInfo.Title = "TimeGladiator API"
	docs.SwaggerInfo.Description = "API of TimeGladiator service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tools := router.Group("/tools")
	auth := router.Group("/auth")
	v1 := router.Group("/api/v1/")
	{

		api.Tools(tools.Group(""))
		api.Auth(auth.Group(""))
		api.User(v1.Group("user"))
		api.TimeEntry(v1.Group("timeEntry"))

	}

	return router
}
