package models

import (
	DBConfig "../Config"
	Utility "../Utility"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	LoginID  string `json："loginId"`
	Password string `json："password"`
	Email    string `json："email"`
	UserName string `json："userName"`
	ToakenID string `json："toakenId"`
}

// TableName テーブル名をreturn
func (b *User) TableName() string {
	return "user"
}

// CreateUser ユーザー追加
func CreateUser(user *User) (err error) {
	Utility.PasswordEncryption(user.Password)
	if err = DBConfig.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByLoginID ログインIDからユーザーを参照
func GetUserByLoginID(user *User) (err error) {
	loginID := user.LoginID
	password := user.Password
	if err = DBConfig.DB.Where("login_id = ? and password = ?", loginID, password).First(user).Error; gorm.IsRecordNotFoundError(err) {
		return err
	}
	return nil
}
