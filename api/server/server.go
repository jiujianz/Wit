package server

import (
	"net/http"

	"../controller"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	router.LoadHTMLFiles("UI/index.html")
	v1.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{})
	})
	{
		v1.POST("/signup", controller.SignUp())
		//server.GET("/signin", controller.SignUp())
	}
	router.Run()
}
