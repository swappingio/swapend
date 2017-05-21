package db

import (
	"fmt"
	"strings"

	"github.com/swappingio/swapend/pkg/mail"
	"github.com/swappingio/swapend/pkg/utils"
	"github.com/swappingio/swapend/pkg/validation"

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
	_, err := validation.ValidateEmail(email)
	if err != nil {
		return fmt.Errorf("Email is not valid.")
	}

	activationcode := utils.GenerateRandomString(100)
	salt := utils.GenerateRandomString(40)
	passHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
	}

	username = strings.ToLower(username)
	email = strings.ToLower(email)

	var userCheck, emailCheck string
	db.QueryRow("SELECT username, email FROM users WHERE username = $1 OR email = $2",
		username, email).Scan(&userCheck, &emailCheck)

	if username == userCheck {
		return fmt.Errorf("Username already exists")
	}

	if email == emailCheck {
		return fmt.Errorf("Email is already registered")
	}

	_, err = db.Exec("INSERT INTO users (username, password, email, salt, activationcode, activated) VALUES ($1, $2, $3, $4, $5, $6)",
		username, passHash, email, salt, activationcode, false)
	if err != nil {
		fmt.Println(err)
	}

	mail.SendActivationEmail(username, email, activationcode)

	return nil

}

func ActivateUser(username string, activationcode string) error {
	username = strings.ToLower(username)

	res, err := db.Exec("UPDATE users SET activated = true, activationcode = '' WHERE username = $1 AND activationcode = $2 AND activated = false",
		username, activationcode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.RowsAffected())
	if res.RowsAffected() != 1 {
		return fmt.Errorf("Invalid activation code")
	}
	return nil
}

func VerifyUser(username string, password string) bool {
	var passHash string
	var salt string

	err := db.QueryRow("SELECT password, salt FROM users WHERE username = $1",
		username).Scan(&passHash, &salt)

	if err != nil {
		fmt.Println("Could not auth user.")
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password+salt))

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
