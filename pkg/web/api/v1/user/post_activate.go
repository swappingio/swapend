package user

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/db"
)

type Activation struct {
	Username       string `form:"username" json:"username" binding:"required"`
	ActivationCode string `form:"activationcode" json:"activationcode" binding:"required"`
}

func postActivate(c *gin.Context) {
	var json Activation
	if c.BindJSON(&json) == nil {
		err := db.ActivateUser(json.Username, json.ActivationCode)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(
				200,
				gin.H{"message": "User " + json.Username + " was activated"},
			)
		}
	}
}
