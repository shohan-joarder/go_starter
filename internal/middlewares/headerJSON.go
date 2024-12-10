package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Middleware to validate JSON payload
func JSONValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next() // Proceed to the next middleware or handler
	}
}
