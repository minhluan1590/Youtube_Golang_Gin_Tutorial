package middlewares

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"user1": "password1",
		"user2": "password2",
	})
}
