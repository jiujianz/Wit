package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	DBConfig "github.com/jiujianz/Wit/Config"
	models "github.com/jiujianz/Wit/Models"
	routes "github.com/jiujianz/Wit/Routes"
)

var err error

func main() {
	DBConfig.DB, err = gorm.Open("mysql",
		DBConfig.DbUrl(DBConfig.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer DBConfig.DB.Close()
	DBConfig.DB.AutoMigrate(&models.User{})

	r := routes.Router()
	r.Run()
}
