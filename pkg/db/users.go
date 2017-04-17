package db

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/coral/swapend/pkg/mail"
	"github.com/goware/emailx"

	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
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

var (
	src = rand.NewSource(time.Now().UnixNano())
)

func CreateUser(username string, password string, email string) error {
	err := emailx.Validate(email)
	if err != nil {
		return fmt.Errorf("Email is not valid.")
	}

	email = emailx.Normalize(email)

	verificationcode := randStringBytesMaskImprSrc(100)
	salt := randStringBytesMaskImprSrc(40)
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

func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
