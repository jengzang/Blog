package middleware

import (
	"blog-backend/config"
	"blog-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.LoadConfig()

		// Get token from cookie
		token, err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"authenticated": false,
				"message":       "No authentication token found",
			})
			c.Abort()
			return
		}

		// Validate token
		claims, err := utils.ValidateToken(token, cfg.JWTSecret)
		if err != nil || !claims.Authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{
				"authenticated": false,
				"message":       "Invalid or expired token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
