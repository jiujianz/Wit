package controller

import (
	"github.com/gin-gonic/gin"
)

func SignUp() {
	var c *gin.Context
	c.JSON(200, gin.H{
		"message": "test",
	})
}
