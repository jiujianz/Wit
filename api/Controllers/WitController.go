package controller

import (
	"fmt"
	"net/http"

	models "github.com/jiujianz/Wit/Models"

	"github.com/gin-gonic/gin"
)

func CreateWit(c *gin.Context) {
	var wit models.Wit
	c.BindJSON(&wit)
	err := models.CreateWit(&wit)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, wit)
	}
}
