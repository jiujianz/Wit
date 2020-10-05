package server

import (
	"../controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup", controller.SignUp())
		//server.GET("/signin", controller.SignUp())
	}
	router.Run()
}
