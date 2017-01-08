package song

import "github.com/gin-gonic/gin"

func getLul(c *gin.Context) {
	c.JSON(200, gin.H{"LUL": "lul"})
}
