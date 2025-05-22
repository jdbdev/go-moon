package env

import (
	"os"
	"strconv"
)

// EnvConfig holds all environment-based configuration settings
// These are loaded once at startup and don't change during runtime
type EnvConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
}

// DatabaseConfig holds database connection settings from environment
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// ServerConfig holds server startup settings from environment
type ServerConfig struct {
	Port         string
	InProduction bool
	UseCache     bool
}

// LoadEnv loads configuration from environment variables with defaults
// This is called once at startup to initialize the application
func LoadEnv() EnvConfig {
	return EnvConfig{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			DBName:   getEnv("POSTGRES_DB", "go_moon"),
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
