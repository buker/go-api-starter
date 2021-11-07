package api

import (
	"github.com/buker/go-api-starter/internal/controller"
	"github.com/gin-gonic/gin"
)

func Tools(g *gin.RouterGroup) {
	g.GET("/health", controller.HealtCheck) // Non-protected
}
