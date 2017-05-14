package song

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/db"
)

func getSong(c *gin.Context) {
	song := c.Param("song")
	test, _ := db.GetSong(song)

	c.JSON(200, test)
}
