package main

import (
	"github.com/coral/swapend/pkg/auth"
	"github.com/coral/swapend/pkg/config"
	"github.com/coral/swapend/pkg/db"
	"github.com/coral/swapend/pkg/session"
	"github.com/coral/swapend/pkg/web/api"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	c config.Specification
)

func init() {

	c = config.GetConfig()
	db.Init()
}

func main() {
	r := gin.Default()

	session.Init(r)

	sAPI := r.Group("/api")
	api.New(sAPI)

	r.Use(static.Serve("/", static.LocalFile("public", true)))

	test := r.Group("/lul")
	test.Use(auth.Auth())
	{
		test.GET("/wut", func(c *gin.Context) {
			c.JSON(400, gin.H{
				"message": "pong",
			})
		})

	}

	r.Run(":4020")
}
