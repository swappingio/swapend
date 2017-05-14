package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/auth"
	"github.com/swappingio/swapend/pkg/config"
	"github.com/swappingio/swapend/pkg/db"
	"github.com/swappingio/swapend/pkg/session"
	"github.com/swappingio/swapend/pkg/web/api"
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
