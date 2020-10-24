package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	DBMS := db_DBMS
	USER := db_USER
	PASS := db_PASS
	PROTOCOL := db_PORT
	DBNAME := db_DBNAME
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	db := gormConnect()

	defer db.Close()
	db.LogMode(true)
}
