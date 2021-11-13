package api

import (
	"github.com/buker/go-api-starter/internal/controllers"
	"github.com/buker/go-api-starter/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// ToolsRoutes defines the routes for the tools api
func TimeEntry(g *gin.RouterGroup) {
	g.Use(middlewares.Authentication())
	g.POST("/", controllers.InsertTimeEntry())

}