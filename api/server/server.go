package server

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	server := gin.Default()
	server.GET("/", IndexRouter)
	server.GET("/signin", SigninFormRoute)
}
