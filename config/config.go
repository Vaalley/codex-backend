package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Loads environment variables from .env file
func LoadConfig() {
	log.Println("🔄 Loading environment variables... 🔑")
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file - make sure it exists and is readable")
	}
	log.Println("✅ Environment variables loaded successfully")
}

// Gets an environment variable by key
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("⚠️  Warning: Environment variable %s is not set", key)
	}
	return value
}

// Returns true if the app is running in production mode
func IsProduction() bool {
	return GetEnv("GO_ENV") == "production"
}
