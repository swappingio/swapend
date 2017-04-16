package api

import (
	"github.com/coral/swapend/pkg/web/api/v1"
	"github.com/gin-gonic/gin"
)

func New(parent *gin.RouterGroup) {
	v1Group := parent.Group("/v1")
	v1.New(v1Group)
}
