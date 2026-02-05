package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	JWTSecret     string
	AccessKeys    []string
	StaticDir     string
}

func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key-change-this"),
		AccessKeys: []string{
			getEnv("ACCESS_KEY_1", "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"), // "secret123" hashed
		},
		StaticDir: getEnv("STATIC_DIR", "../blog/public"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
