package user

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/auth"
)

func New(parent *gin.RouterGroup) {

	parent.POST("/create", postCreate)
	parent.POST("/activate", postActivate)
	parent.POST("/login", postLogin)
	parent.POST("/logout", postLogout)

	authorized := parent.Group("/")
	authorized.Use(auth.Auth())
	{
		authorized.GET("/auth", getAuth)
	}
}
