package Utility

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func PasswordEncryption(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	fmt.Println(string(hash))
}