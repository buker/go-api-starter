package api

import (
	"github.com/buker/go-api-starter/internal/controller"
	"github.com/buker/go-api-starter/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Users(g *gin.RouterGroup) {
	g.GET("", controller.GetUsers)    // Non-protected
	g.GET("/:id", controller.GetUser) // Non-protected
	g.Use(middleware.JWT())
	g.POST("", controller.CreateUser) // Protected
	g.PUT("", controller.UpdateUser)  // Protected
}
