package sanity

import "github.com/gin-gonic/gin"

func getSanity(c *gin.Context) {
	c.JSON(200, gin.H{"IT IS": "OK"})
}
