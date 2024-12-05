package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shohan-joarder/go_pos/internal/services"
)

type AuthorizationMiddleware struct {
	rolePermissionService *services.RolePermissionService
}

// Constructor
func NewAuthorizationMiddleware(rolePermissionService *services.RolePermissionService) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{rolePermissionService: rolePermissionService}
}

// Middleware Handler
func (m *AuthorizationMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract role ID from context
		roleID, exists := c.Get("role_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role ID is missing"})
			c.Abort()
			return
		}

		// Ensure roleID is a valid type (float64 is often used when unmarshalling JSON)
		parsedRoleID, ok := roleID.(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role ID format"})
			c.Abort()
			return
		}

		// Extract the route pattern and HTTP method
		requestURL := c.FullPath()        // Gets the route pattern (e.g., "/products/:id")
		requestMethod := c.Request.Method // Extracts the HTTP method (e.g., "POST")

		// Check permissions via the service
		isAllowed, err := m.rolePermissionService.FilterPermissionByRoleIdAndURLMethod(uint(parsedRoleID), requestURL, requestMethod)
		if err != nil {
			// Log the error for debugging purposes
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Permission check failed: %v", err)})
			c.Abort()
			return
		}

		if !isAllowed {
			// Permission denied
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: insufficient permissions"})
			c.Abort()
			return
		}

		// Proceed to the next middleware/handler
		c.Next()
	}
}

// Helper function to match dynamic URLs
// func matchDynamicRoute(requestURL, requestMethod string, rules []AccessRule) bool {
// 	for _, rule := range rules {
// 		// Convert dynamic segments (e.g., ":product_id") in the rule's URL to regex patterns
// 		pattern := "^" + regexp.MustCompile(`:[^/]+`).ReplaceAllString(rule.URL, `[^/]+`) + "$"

// 		// Check if the request matches the rule's URL and HTTP method
// 		matched, _ := regexp.MatchString(pattern, requestURL)
// 		if matched && strings.EqualFold(requestMethod, rule.Method) {
// 			return true
// 		}
// 	}
// 	return false
// }
