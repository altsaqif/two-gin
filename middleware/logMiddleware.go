package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		userAgent := c.Request.UserAgent()
		path := c.Request.URL.Path

		log.Printf("[LOG] %s - [%v] \"%s %s %d %v \"%s\"\n", clientIP, t, method, path, statusCode, latency, userAgent)
	}
}
