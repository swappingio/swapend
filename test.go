package main

import (
	"log"
	"time"

	"github.com/coral/swapend/db"
	"github.com/pajlada/kkonfig"
)

type Specification struct {
	Database db.DatabaseSettings
}

func init() {

	var s Specification
	var test []string
	test = append(test, "konfig.json")
	err := kkonfig.Process("web", test, &s)
	if err != nil {
		log.Fatal(err)
	}
	db.Init(s.Database)
}

func main() {
	createdUUID, _ := db.CreateSong(db.Song{Created: time.Now(), Name: "Gachitest"})
	db.CreateVersion(createdUUID)
}
