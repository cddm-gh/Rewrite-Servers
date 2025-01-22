package config

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	OREServiceURL string
	Port          int
	// We can add more config fields here as needed
}

var cfg *Config

// Initialize loads all configuration values from environment variables
func Initialize() *Config {
	if cfg != nil {
		return cfg
	}

	cfg = &Config{
		OREServiceURL: getEnvWithDefault("ORE_SERVICE_URL", "http://localhost:8625/api"),
		Port:          getEnvAsIntWithDefault("PORT", 8628),
	}

	// Log the configuration (excluding sensitive values)
	slog.Info("Configuration loaded",
		"ore_service_url", cfg.OREServiceURL,
		"port", cfg.Port,
	)

	return cfg
}

// Get returns the current configuration
func Get() *Config {
	if cfg == nil {
		return Initialize()
	}
	return cfg
}

// Helper function to get environment variable with default value
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Helper function to get environment variable as int with default value
func getEnvAsIntWithDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		slog.Warn("Invalid integer value in environment variable", "key", key, "value", value)
	}
	return defaultValue
}
