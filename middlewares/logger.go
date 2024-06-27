package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	log.SetFlags(0)
	log.SetOutput(gin.DefaultWriter)
	log.SetFlags(log.Ldate | log.Ltime)
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		log.Printf("Request: %s %s, Latency: %v", c.Request.Method, c.Request.URL.Path, latency.Round(time.Microsecond))
	}
}
