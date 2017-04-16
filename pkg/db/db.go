package db

import (
	"fmt"
	"log"

	"github.com/coral/swapend/pkg/config"
	"github.com/jackc/pgx"
)

var db *pgx.Conn

func Init() {
	c := config.GetConfig().Database
	dbInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		c.Username, c.Password, c.Hostname, c.Dbname)

	var err error
	dbConfig, err := pgx.ParseConnectionString(dbInfo)
	db, err = pgx.Connect(dbConfig)
	if err != nil {
		log.Fatal(err)
		fmt.Println("TEST")
	}
}
