package handlers

import (
	"blog-backend/config"
	"blog-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Key string `json:"key" binding:"required"`
}

// Authenticate handles the authentication request
func Authenticate(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request",
		})
		return
	}

	cfg := config.LoadConfig()

	// Check if the provided key matches any of the stored hashed keys
	authenticated := false
	for _, hashedKey := range cfg.AccessKeys {
		if utils.CheckPassword(req.Key, hashedKey) {
			authenticated = true
			break
		}
	}

	if !authenticated {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid access key",
		})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(cfg.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to generate token",
		})
		return
	}

	// Set cookie with token
	c.SetCookie(
		"auth_token",
		token,
		86400,    // 24 hours
		"/",
		"",
		false,    // Set to true if using HTTPS
		true,     // HttpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Authentication successful",
	})
}
