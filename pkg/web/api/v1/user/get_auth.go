package user

import "github.com/gin-gonic/gin"

func getAuth(c *gin.Context) {
	c.JSON(200, gin.H{"YES": "YES"})
}
