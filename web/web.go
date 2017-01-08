package main

import (
	"log"

	"github.com/coral/swapend/db"
	"github.com/coral/swapend/web/api"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pajlada/kkonfig"
)

type Specification struct {
	Database db.DatabaseSettings
}

func init() {

	var s Specification
	err := kkonfig.Process("web", "konfig.json", &s)
	if err != nil {
		log.Fatal(err)
	}
	db.Init(s.Database)
}

func main() {
	r := gin.Default()

	sAPI := r.Group("/api")
	api.New(sAPI)

	r.Use(static.Serve("/", static.LocalFile("public", true)))

	r.Run(":4020")
}
