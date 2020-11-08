package models

import (
	DBConfig "../Config"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	LoginID  string `json："LoginId"`
	Password string `json："password"`
	Email    string `json："email"`
	UserName string `json："userName"`
}

func (b *User) TableName() string {
	return "user"
}

func CreateUser(user *User) (err error) {
	if err = DBConfig.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByLoginID(user *User) (err error) {
	loginID := user.LoginID
	password := user.Password
	if err = DBConfig.DB.Where("login_id = ? and password = ?", loginID, password).First(user).Error; err != nil {
		return err
	}
	return nil
}
