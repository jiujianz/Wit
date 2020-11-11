package service

import (
	"fmt"

	models "../Models"

	"github.com/gin-gonic/gin"
)

func GetUserByLoginID(c *gin.Context) models.User {
	var user models.User
	c.BindJSON(&user)
	err := models.GetUserByLoginID(&user)
	if err != nil {
		//c.AbortWithStatus(http.StatusNotFound)
		// TODO error response
		return ""
	} else {
		fmt.Println(user)
		return user
	}
}

func CreateToaken(c *gin.Context) models.Toaken {
	var Toaken models.Toaken
	c.BindJSON(&Toaken)
	err := models.CreateToaken(&Toaken)
	if err != nil {
		fmt.Println(err.Error())
		// TODO error response
		return ""
	} else {
		return Toaken
	}
}
