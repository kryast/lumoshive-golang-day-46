package middleware

import (
	"day-46/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPWhitelist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !config.AllowedIPs[ip] {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
