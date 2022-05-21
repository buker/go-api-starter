package api

import (
	"github.com/buker/TimeGladiator/internal/controllers"
	"github.com/gin-gonic/gin"
)

// ToolsRoutes defines the routes for the tools api
func Auth(g *gin.RouterGroup) {
	g.POST("/users/signup", controllers.SignUp())
	g.POST("/users/login", controllers.Login())
}
