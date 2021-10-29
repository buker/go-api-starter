package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run()
}

// The engine with all endpoints is now extracted from the main function
func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthEndpoint)

	return r
}

func healthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
