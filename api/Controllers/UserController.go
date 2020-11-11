package controller

import (
	"fmt"
	"net/http"

	models "../Models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
