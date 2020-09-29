package server

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engin {
	r := gin.Default()

	user := r.Group("/dev/user/") {
		control := controller.UserController{}
		control.POST("", control.Create)
	}
}