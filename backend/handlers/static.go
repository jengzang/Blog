package handlers

import (
	"blog-backend/config"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// ServeStatic serves static files from the Hexo public directory
func ServeStatic(c *gin.Context) {
	cfg := config.LoadConfig()
	path := c.Request.URL.Path

	// Default to index.html for root path
	if path == "/" {
		path = "/index.html"
	}

	// Add .html extension if not present and file doesn't exist
	filePath := filepath.Join(cfg.StaticDir, path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if !strings.HasSuffix(path, ".html") && !strings.Contains(path, ".") {
			htmlPath := filepath.Join(cfg.StaticDir, path, "index.html")
			if _, err := os.Stat(htmlPath); err == nil {
				filePath = htmlPath
			}
		}
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Try 404.html
		notFoundPath := filepath.Join(cfg.StaticDir, "404.html")
		if _, err := os.Stat(notFoundPath); err == nil {
			c.File(notFoundPath)
			return
		}
		c.String(http.StatusNotFound, "404 - Page Not Found")
		return
	}

	c.File(filePath)
}
