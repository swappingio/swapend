package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/web/api/v1"
)

func New(parent *gin.RouterGroup) {
	v1Group := parent.Group("/v1")
	v1.New(v1Group)
}
