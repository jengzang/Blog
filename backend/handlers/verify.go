package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VerifyAuth checks if the user is authenticated
func VerifyAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"message":       "User is authenticated",
	})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// Clear the auth cookie
	c.SetCookie(
		"auth_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged out successfully",
	})
}
