package user

import "github.com/gin-gonic/gin"

func New(parent *gin.RouterGroup) {

	parent.POST("/create", postCreate)
	parent.POST("/activate", postActivate)
}
