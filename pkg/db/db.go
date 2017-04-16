package db

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

type DatabaseSettings struct {
	Username string `default:"coral"`
	Password string `default:"lolwut"`
	Hostname string `default:"localhost"`
	Dbname   string `default:"lolwut"`
}

var db *pgx.Conn

func Init(c DatabaseSettings) {
	dbInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		c.Username, c.Password, c.Hostname, c.Dbname)

	var err error
	dbConfig, err := pgx.ParseConnectionString(dbInfo)
	db, err = pgx.Connect(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
}
