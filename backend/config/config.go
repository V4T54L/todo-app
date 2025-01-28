package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the application configuration
type Config struct {
	ServerPort   string
	DatabaseURI  string
	MaxOpenConns int
	MaxIdleConns int
}

var configInstance *Config

// GetConfig returns a new instance of Config.
// It loads configuration from a .env file at the specified path.
func GetConfig() (*Config, error) {
	if configInstance == nil {
		instance := &Config{}

		serverPort := "8080" // Default value for server port
		if val, err := getStr("SERVER_PORT", &serverPort); err == nil {
			instance.ServerPort = val
		}

		if val, err := getStr("DB_URI", nil); err == nil {
			instance.DatabaseURI = val
		} else {
			return configInstance, fmt.Errorf("missing required environment variable DB_URI: %w", err)
		}

		if val, err := getInt("MAX_IDLE_CONNS", nil); err == nil {
			instance.MaxIdleConns = val
		} else {
			return nil, fmt.Errorf("missing required environment variable MAX_IDLE_CONNS: %w", err)
		}

		if val, err := getInt("MAX_OPEN_CONNS", nil); err == nil {
			instance.MaxOpenConns = val
		} else {
			return nil, fmt.Errorf("missing required environment variable MAX_OPEN_CONNS: %w", err)
		}

		configInstance = instance
	}
	return configInstance, nil
}

// LoadConfigurationFile loads configuration from a .env file at the specified path.
func LoadConfigurationFile(filePath string) error {
	if err := godotenv.Load(filePath); err != nil {
		return fmt.Errorf("error loading .env file '%s': %v", filePath, err)
	}
	return nil
}

// getStr retrieves an environment variable by key; returns fallback if not found.
// retuns error if missing environment variable and fallback is nil.
func getStr(key string, fallback *string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		if fallback == nil {
			return "", fmt.Errorf("missing required environment variable: %s", key)
		} else {
			return *fallback, nil
		}
	}
	return value, nil
}

// getInt retrieves an environment variable by key; returns fallback if not found.
// retuns error if missing environment variable and fallback is nil.
func getInt(key string, fallback *int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		if fallback == nil {
			return 0, fmt.Errorf("missing required environment variable: %s", key)
		} else {
			return *fallback, nil
		}
	}

	return strconv.Atoi(value)
}
