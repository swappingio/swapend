package v1

import (
	"github.com/coral/swapend/web/api/v1/song"
	"github.com/gin-gonic/gin"
)

func New(parent *gin.RouterGroup) {
	songGroup := parent.Group("/song")
	song.New(songGroup)
}
