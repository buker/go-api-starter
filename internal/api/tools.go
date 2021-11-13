package api

import (
	"github.com/buker/go-api-starter/internal/controllers"
	"github.com/gin-gonic/gin"
)

// ToolsRoutes defines the routes for the tools api
func Tools(g *gin.RouterGroup) {
	g.GET("/health", controllers.HealtCheck) // Non-protected
}
