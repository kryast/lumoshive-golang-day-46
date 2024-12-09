package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var failedAttempts = make(map[string]int)
var mu sync.Mutex

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		if failedAttempts[ip] >= 3 {
			mu.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many attempts, try again later"})
			c.Abort()
			return
		}

		failedAttempts[ip]++
		mu.Unlock()

		c.Next()

		time.AfterFunc(1*time.Minute, func() {
			mu.Lock()
			failedAttempts[ip]--
			mu.Unlock()
		})
	}
}
