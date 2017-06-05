package user

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/db"
)

type Update struct {
	Username    string `form:"username" json:"username" binding:"required"`
	OldPassword string `form:"oldpassword" json:"oldpassword" binding:"required"`
	NewPassword string `form:"newpassword" json:"newpassword"`
	Email       string `form:"email" json:"email"`
}

func postUpdate(c *gin.Context) {
	var json Update

	if c.BindJSON(&json) == nil {

		userid, verified := db.VerifyUser(json.Username, json.OldPassword)

		if json.Username != "" {
			//Update Username
		}

		if json.NewPassword != "" && verified {
			db.SetPassword(userid, json.NewPassword)
		}
	}
}
