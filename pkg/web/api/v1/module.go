package v1

import (
	"github.com/coral/swapend/pkg/web/api/v1/sanity"
	"github.com/coral/swapend/pkg/web/api/v1/song"
	"github.com/gin-gonic/gin"
)

func New(parent *gin.RouterGroup) {
	songGroup := parent.Group("/song")
	song.New(songGroup)
	sanityGroup := parent.Group("/sanity")
	sanity.New(sanityGroup)
}
