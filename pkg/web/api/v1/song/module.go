package song

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/auth"
)

func New(parent *gin.RouterGroup) {

	parent.GET("/get/:song", getSong)

	authorized := parent.Group("/")
	authorized.Use(auth.Auth())
	{
		authorized.GET("/test/:song", getLOL)
		authorized.POST("/upload", postVersion)
		authorized.POST("/create", postSong)
	}
}
