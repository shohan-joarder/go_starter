package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Replace with a secure secret

// AuthMiddleware validates the JWT token in the Authorization header.
func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the Authorization header
		// Retrieve the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			respondUnauthorized(c, "Missing Authorization header")
			return
		}

		// Remove "Bearer " prefix if present
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			respondUnauthorized(c, "Invalid or expired token")
			return
		}

		// Decode token and extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			respondUnauthorized(c, "Invalid token claims")
			return
		}

		// fmt.Fprintf(os.Stderr, "Claims: %+v\n", claims)

		// Extract and validate required claims
		email, emailExists := claims["email"].(string)
		userID, userIDExists := claims["id"]
		roleID, roleIDExists := claims["role"]

		// fmt.Fprintf(os.Stderr, "Email: %s, User ID: %v, Role ID: %v\n", email, userID, roleID)

		// Check if required claims are present

		if !emailExists || !userIDExists || !roleIDExists {
			respondUnauthorized(c, "Required claims missing in token")
			return
		}

		// Add claims to the context
		c.Set("email", email)
		c.Set("role_id", roleID)
		c.Set("user_id", userID)

		// Proceed to the next middleware/handler
		c.Next()
	}
}

// respondUnauthorized sends an unauthorized response and aborts the context
func respondUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}
