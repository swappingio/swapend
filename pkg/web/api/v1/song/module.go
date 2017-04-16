package song

import "github.com/gin-gonic/gin"

func New(parent *gin.RouterGroup) {

	parent.GET("/get/:song", getSong)

	parent.POST("/upload", postVersion)
	parent.POST("/create", postSong)
}
