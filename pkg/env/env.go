package env

import (
	"os"
	"strconv"
)

// Config holds all startup configuration settings
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

// DatabaseConfig holds database-specific configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port         string
	InProduction bool
	UseCache     bool
}

// LoadEnv loads configuration from environment variables with defaults
func LoadEnv() Config {
	return Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "go_moon_db"),
		},
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			InProduction: getEnvBool("IN_PRODUCTION", false),
			UseCache:     getEnvBool("USE_CACHE", true),
		},
	}
}

// Helper functions to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			return boolValue
		}
	}
	return defaultValue
}
