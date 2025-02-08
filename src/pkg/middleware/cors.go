package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, Change as needed
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Access-Control-Max-Age: 0 (or not setting it at all) ensures browsers don’t cache preflight responses.
		// Cache-Control: no-store, no-cache, must-revalidate, max-age=0 prevents caching on both server and client sides.
		// Pragma: no-cache is for older HTTP/1.0 caches.
		// Expires: 0 makes sure caching expires immediately.
		c.Writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
