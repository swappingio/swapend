package db

import (
	"fmt"
	"strings"

	"github.com/swappingio/swapend/pkg/auth"
	"github.com/swappingio/swapend/pkg/mail"
	"github.com/swappingio/swapend/pkg/utils"
	"github.com/swappingio/swapend/pkg/validation"
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

	var newID int64

	err = db.QueryRow("INSERT INTO users (username, password, email, salt, activationcode, activated) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		username, "placeholder", email, "saltholder", activationcode, false).Scan(&newID)
	if err != nil {
		fmt.Println(err)
	}

	SetPassword(newID, password)

	mail.SendActivationEmail(username, email, activationcode)

	return nil

}

func SetPassword(userid int64, password string) {
	passHash, salt := auth.CreatePassword(password)
	_, err := db.Exec("UPDATE users SET password = $1, salt = $2 WHERE id = $3", passHash, salt, userid)
	if err != nil {
		fmt.Println("error on setting password")
		fmt.Println(err)
	}
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

func VerifyUser(username string, password string) (int64, bool) {
	var passHash string
	var salt string
	var id int64
	username = strings.ToLower(username)

	err := db.QueryRow("SELECT password, salt, id FROM users WHERE username = $1",
		username).Scan(&passHash, &salt, &id)

	if err != nil {
		fmt.Println("Could not auth user.")
		return 0, false
	}

	matches := auth.VerifyPassword(password, passHash, salt)

	if matches {
		return id, true
	} else {
		return 0, false
	}
}
