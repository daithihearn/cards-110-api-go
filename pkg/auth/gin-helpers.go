package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckValidated returns the user ID from the Gin context if it exists.
// If the user ID is not found the gin context is aborted with a 401 status code.
// Returns a bool indicating if the user ID was found and the user ID.
func CheckValidated(c *gin.Context) (string, bool) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token is missing or has not been validated"})
		return "", false
	}

	claims := user.(*CustomClaims)
	return claims.RegisteredClaims.Subject, true
}
