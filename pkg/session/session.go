package session

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/swappingio/swapend/pkg/config"
)

func Init(r *gin.Engine) {
	c := config.GetConfig()

	store, _ := sessions.NewRedisStore(c.Redis.Timeout,
		"tcp",
		c.Redis.Hostname+":"+strconv.Itoa(c.Redis.Port),
		"",
		[]byte(c.Sessions.CookieSecret),
	)
	r.Use(sessions.Sessions(c.Sessions.StorageName, store))

	fmt.Println("Connected to Redis")
}
