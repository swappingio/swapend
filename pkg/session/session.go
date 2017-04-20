package session

import (
	"github.com/coral/swapend/pkg/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	c := config.GetConfig()
	store, _ := sessions.NewRedisStore(c.Redis.Timeout, "tcp", "localhost:6379", "",
		[]byte(c.Sessions.CookieSecret))
	r.Use(sessions.Sessions(c.Sessions.StorageName, store))

}
