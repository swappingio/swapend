package auth

import (
	"fmt"

	"github.com/swappingio/swapend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(password string) (passHash string, salt string) {
	newSalt := utils.GenerateRandomString(40)
	newPassHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
	}

	return string(newPassHash), newSalt
}

func VerifyPassword(password string, passHash string, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password+salt))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
