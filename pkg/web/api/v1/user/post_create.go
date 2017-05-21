package user

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/db"
)

type NewUser struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

func postCreate(c *gin.Context) {
	var json NewUser
	if c.BindJSON(&json) == nil {
		err := db.CreateUser(json.Username, json.Password, json.Email)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(
				200,
				gin.H{"message": "User " + json.Username + " was created"},
			)
		}
	}
}
