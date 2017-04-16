package song

import (
	"github.com/coral/swapend/pkg/db"
	"github.com/gin-gonic/gin"
)

func getSong(c *gin.Context) {
	song := c.Param("song")
	test, _ := db.GetSong(song)

	c.JSON(200, test)
}
