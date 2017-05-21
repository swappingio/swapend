package user

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/auth"
	"github.com/swappingio/swapend/pkg/db"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func postLogin(c *gin.Context) {
	var json Login
	if c.BindJSON(&json) == nil {
		verified := db.VerifyUser(json.Username, json.Password)
		if verified {
			auth.SetAuth(1, c)
		}
	}
}
