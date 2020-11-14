package controller

import (
	"fmt"
	"net/http"

	service "../Services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//var jwtService service.JWTService = service.JWTAuthService()
	//var login LoginInterface = LoginHandler(jwtService)
	user := service.GetUserByLoginID(c)
	toaken := service.GenerateToaken(user.LoginID, true)
	fmt.Println(toaken)
	c.JSON(http.StatusOK, user)
}
