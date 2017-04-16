package db

type User struct {
	Id       int64
	Username string
	Password string
	Hash     string
}
