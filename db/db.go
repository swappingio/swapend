package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseSettings struct {
	Username string `default:"coral"`
	Password string `default:"lolwut"`
	Host     string `default:"localhost"`
	Dbname   string `default:"lolwut"`
}

var db *sql.DB

func Init(c DatabaseSettings) {
	dbInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=require", c.Username, c.Password, c.Host, c.Dbname)

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
