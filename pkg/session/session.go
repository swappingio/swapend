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

	store, err := sessions.NewRedisStore(c.Redis.Timeout,
		"tcp",
		c.Redis.Hostname+":"+strconv.Itoa(c.Redis.Port),
		"",
		[]byte(c.Sessions.CookieSecret),
	)
	r.Use(sessions.Sessions(c.Sessions.StorageName, store))
	
	if err == nil {
		fmt.Println("Connected to Redis")
	} else {
		fmt.Println("Couldn't connect to Redis")
	}
}
