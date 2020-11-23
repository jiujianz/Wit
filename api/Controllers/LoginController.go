package controller

import (
	"net/http"

	service "../Services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user, err := service.GetUserByLoginID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "data not found")
		return
	}
	toaken := service.GenerateToaken(user.LoginID, true)
	user.ToakenID = toaken
	c.JSON(http.StatusOK, user)
}
