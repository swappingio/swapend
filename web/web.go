package main

import (
	"github.com/coral/swapend/web/api"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sAPI := r.Group("/api")
	api.New(sAPI)

	r.Use(static.Serve("/", static.LocalFile("public", true)))

	r.Run(":4020")
}
