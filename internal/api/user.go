package api

import (
	"github.com/buker/TimeGladiator/internal/controllers"
	"github.com/buker/TimeGladiator/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// ToolsRoutes defines the routes for the tools api
func User(g *gin.RouterGroup) {
	g.Use(middlewares.Authentication())
	g.GET("/users", controllers.GetUsers())
	g.GET("/users/:user_id", controllers.GetUser())
}
