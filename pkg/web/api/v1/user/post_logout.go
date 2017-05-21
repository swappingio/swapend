package user

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/auth"
)

func postLogout(c *gin.Context) {
	auth.KillAuth(c)
	c.JSON(200, gin.H{"message": "Goodbye"})
}
