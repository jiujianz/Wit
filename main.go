package main

import (
	"fmt"

	DBConfig "./api/Config"
	models "./api/Models"
	routes "./api/Routes"
	"github.com/jinzhu/gorm"
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
