package models

import (
	DBConfig "../Config"
	_ "github.com/go-sql-driver/mysql"
)

type Wit struct {
	Content string `json: "content"`
	Image   string `json: "image"`
	Video   string `json: "video"`
	UserID  int    `json："userId"`
}

func (b *Wit) TableName() string {
	return "wit"
}

func CreateWit(wit *Wit) (err error) {
	if err = DBConfig.DB.Create(wit).Error; err != nil {
		return err
	}
	return nil
}
