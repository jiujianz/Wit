package controller

import (
	"fmt"
	"net/http"

	models "../Models"

	"github.com/gin-gonic/gin"
)

func CreateToaken(c *gin.Context) {
	var Toaken models.Toaken
	c.BindJSON(&Toaken)
	err := models.CreateToaken(&Toaken)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Toaken)
	}
}
