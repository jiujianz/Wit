package routes

import (
	"net/http"

	controller "../Controllers"
	"github.com/gin-gonic/gin"
)

// Router setupRouter
func Router() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		router.LoadHTMLFiles("UI/index.html")
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", map[string]interface{}{})
		})

		v1.POST("/signup", controller.CreateUser)
	}

	return router
}
