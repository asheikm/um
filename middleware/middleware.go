package middleware

import (
	"net/http"

	"fmt"
	"um/models"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func Cors() {}

func JWTAuth() {}

func getUserFromSession(c *gin.Context) (*models.User, error) {
	// Replace with your logic to retrieve user from session
	// (e.g., read session ID from cookie, decode session data, extract user)
	return nil, fmt.Errorf("Unable to retrieve user from session")
}

func RBACMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user information from session or JWT (replace with your method)
		user, err := getUserFromSession(c)
		if err != nil {
			// Handle error (e.g., unauthorized access)
			return
		}

		// Extract path and method from request
		path := c.Request.URL.Path
		method := c.Request.Method

		// Check if user has permission for the requested action and resource
		hasPermission := enforcer.Enforce(user.Role, path, method)

		if !hasPermission {
			// Handle unauthorized access (e.g., return 403 Forbidden)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// User has permission, proceed with the request
		c.Next()
	}
}
