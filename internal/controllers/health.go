package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health check
// @Description Get health status
// @Produce  json
// @Success 200
// @Router /tools/health [get]
func HealtCheck(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": "ok"})

}
