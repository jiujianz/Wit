package models

import (
	"time"

	DBConfig "../Config"
	_ "github.com/go-sql-driver/mysql"
)

type Toaken struct {
	ToakenID  string    `json:"toakenId"`
	UserID    int       `json:"userId"`
	limitTime time.Time `json:"limitTime"`
}

func (b *Toaken) TableName() string {
	return "toaken_id"
}

func CreateToaken(token *Toaken) (err error) {
	if err = DBConfig.DB.Create(token).Error; err != nil {
		return err
	}
	return nil
}
