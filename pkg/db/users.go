package db

import (
	"fmt"

	"github.com/coral/swapend/pkg/mail"
	"github.com/coral/swapend/pkg/utils"
	"github.com/coral/swapend/pkg/validation"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int64
	Username     string
	Password     string
	Email        string
	Salt         string
	Reset        string
	Verified     bool
	Verification string
}

func CreateUser(username string, password string, email string) error {
	err := validation.ValidateEmail(email)
	if err != nil {
		return fmt.Errorf("Email is not valid.")
	}

	verificationcode := utils.GenerateRandomString(100)
	salt := GenerateRandomString(40)
	passHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
	}

	_, err = db.Exec("INSERT INTO users (username, password, email, salt, verificationcode, verified) VALUES ($1, $2, $3, $4, $5, $6)",
		username, passHash, email, salt, verificationcode, false)
	if err != nil {
		fmt.Println(err)
	}

	mail.SendVerificationEmail(username, email, "AJIOGJIEOGJGO")

	return nil

}

func VerifyUser(username string, password string) bool {
	var passHash string
	var salt string

	err = db.QueryRow("SELECT password, salt FROM users WHERE username = $1",
		username).Scan(&passHash, &salt)

	if err != nil {
		fmt.Println("Could not auth user.")
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password+salt))

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
