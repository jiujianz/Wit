package controller

import (
	"fmt"
	"net/http"

	service "../Services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user := service.GetUserByLoginID(c)
	toaken := service.GenerateToaken(user.LoginID, true)
	user.ToakenID = toaken
	fmt.Println(user)
	c.JSON(http.StatusOK, user)
}
