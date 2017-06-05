package song

import "github.com/gin-gonic/gin"

func getLOL(c *gin.Context) {
	c.JSON(200, "hello")
}
