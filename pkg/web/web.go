package main

import (
	"github.com/coral/swapend/pkg/config"
	"github.com/coral/swapend/pkg/db"
	"github.com/coral/swapend/pkg/web/api"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {

	s := config.GetConfig()
	db.Init(s.Database)
}

func main() {
	r := gin.Default()

	sAPI := r.Group("/api")
	api.New(sAPI)

	r.Use(static.Serve("/", static.LocalFile("public", true)))

	r.Run(":4020")
}
