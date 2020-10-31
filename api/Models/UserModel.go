package models

import (
	DBConfig "../Config"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserID   int    `json："userId"`
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

// func insertUser(createTime time.Time, email string, password string, userID int, name string) {
// 	db := db.Connect()
// 	defer db.close()

// 	db.Create(&User{
// 		CreatedTime: createTime,
// 		Email:       email,
// 		Password:    password,
// 		UserID:      userID,
// 		UserName:    name,
// 	})
// }

// func createUser(c *gin.Context) {
// 	createTime := c.PostForm("CreateedTime")
// 	email := c.PostForm("CreateedTime")
// 	password := c.PostForm("Password")
// 	id := c.PostForm("Id")
// 	UserName := c.PostForm("UserName")
// 	insertUser(createTime, email, password, id, UserName)
// 	c.Redirect(302, "/")
// }
