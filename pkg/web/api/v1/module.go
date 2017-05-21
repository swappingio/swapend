package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/web/api/v1/sanity"
	"github.com/swappingio/swapend/pkg/web/api/v1/song"
	"github.com/swappingio/swapend/pkg/web/api/v1/user"
)

func New(parent *gin.RouterGroup) {
	songGroup := parent.Group("/song")
	song.New(songGroup)
	sanityGroup := parent.Group("/sanity")
	sanity.New(sanityGroup)
	userGroup := parent.Group("/user")
	user.New(userGroup)
}
