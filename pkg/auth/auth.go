package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		s := session.Get("authenticated")
		if s != nil {
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"message": "NOT AUTHENTICATED",
			})
			c.Done()
		}

	}
}

func IsAuthenticated() bool {
	session := sessions.Default(c)
	s := session.Get("authenticated")
	return s
}
