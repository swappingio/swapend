package song

import "github.com/gin-gonic/gin"

func New(parent *gin.RouterGroup) {
	parent.POST("/upload", postSong)
	parent.GET("/test", getLul)
}
