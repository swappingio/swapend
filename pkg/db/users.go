package db

import (
	"math/rand"
	"time"

	"github.com/coral/swapend/pkg/mail"

	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
	Salt     string
	Reset    string
	Verified bool
}

var (
	src = rand.NewSource(time.Now().UnixNano())
)

func CreateUser(username string, password string, email string) {
	salt := randStringBytesMaskImprSrc(40)
	passHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {

	}

	_, err = db.Exec("INSERT INTO users (username, password, email, salt, verified) VALUES ($1, $2, $3, $4, FALSE)",
		username, passHash, email, salt)
	if err != nil {

	}

	go mail.SendVerificationEmail(username, email, "AJIOGJIEOGJGO")

}

func VerifyUser(username string, password string) bool {
	var passHash string
	var salt string

	db.QueryRow("SELECT password, salt FROM users WHERE username = $1",
		username).Scan(&passHash, &salt)
	err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password+salt))

	if err != nil {
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
