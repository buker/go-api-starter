package controller

import "github.com/gin-gonic/gin"

func render(c *gin.Context, data interface{}, statusCode int) {
	if statusCode >= 400 {
		c.AbortWithStatusJSON(statusCode, data)
		return
	}

	c.JSON(statusCode, data)
}
