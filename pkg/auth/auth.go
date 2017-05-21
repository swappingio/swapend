package auth

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		s := session.Get("authenticated")
		if s != nil && s == true {
			fmt.Println(s)
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"message": "NOT AUTHENTICATED",
			})
			c.Abort()
		}

	}
}

func SetAuth(status int, c *gin.Context) {
	sessions := sessions.Default(c)
	sessions.Set("authenticated", true)
	sessions.Set("status", status)
	sessions.Save()
}

func KillAuth(c *gin.Context) {
	sessions := sessions.Default(c)
	sessions.Clear()
	sessions.Save()
}

/*
func IsAuthenticated() bool {
	session := sessions.Default(c)
	s := session.Get("authenticated")
	return s
}
*/
