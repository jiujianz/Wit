package controller

import (
	"net/http"

	service "../Services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user := service.GetUserByLoginID(c)
	c.JSON(http.StatusOK, user)
}
