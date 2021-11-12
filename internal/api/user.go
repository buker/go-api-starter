package api

import (
	"github.com/buker/go-api-starter/internal/controllers"
	"github.com/buker/go-api-starter/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// ToolsRoutes defines the routes for the tools api
func User(g *gin.RouterGroup) {
	g.Use(middlewares.Authentication())
	g.GET("/users", controllers.GetUsers())
	g.GET("/users/:user_id", controllers.GetUser())
}
