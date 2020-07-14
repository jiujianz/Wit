package db

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := <DB user>
	PASS := <DB Password>
	PROTOCOL := tcp(<DBのIPアドレス>:<PORT>)
	DBNAME := <DB NAME>
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