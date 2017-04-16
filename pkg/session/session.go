package session

import (
	"github.com/coral/swapend/pkg/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	c := config.GetConfig().Authentication
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "",
		[]byte(c.CookieSecret))
	r.Use(sessions.Sessions(c.StorageName, store))

}
