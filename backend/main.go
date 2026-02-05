package main

import (
	"blog-backend/config"
	"blog-backend/handlers"
	"blog-backend/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/auth", handlers.Authenticate)
		api.GET("/verify", middleware.AuthMiddleware(), handlers.VerifyAuth)
		api.POST("/logout", handlers.Logout)
	}

	// Serve static files from Hexo public directory
	r.Static("/css", "../blog/public/css")
	r.Static("/js", "../blog/public/js")
	r.Static("/img", "../blog/public/img")
	r.Static("/fonts", "../blog/public/fonts")

	// Serve HTML files
	r.NoRoute(handlers.ServeStatic)

	// Start server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
