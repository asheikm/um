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

		path := c.Request.URL.Path
		method := c.Request.Method

		hasPermission := enforcer.Enforce(user.Role, path, method)

		if !hasPermission {
			// Handle unauthorized access (e.g., return 403 Forbidden)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
