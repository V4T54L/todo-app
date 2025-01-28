package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// Temporary environment variable keys for testing
	testServerPortKey   = "SERVER_PORT"
	testDbURIKey        = "DB_URI"
	testMaxIdleConnsKey = "MAX_IDLE_CONNS"
	testMaxOpenConnsKey = "MAX_OPEN_CONNS"
)

// Setup function to clean up after tests
func setup(t *testing.T) {
	// Set temporary environment variables before testing
	t.Setenv(testServerPortKey, "9090")
	t.Setenv(testDbURIKey, "postgres://user:password@localhost:5432/mydb")
	t.Setenv(testMaxIdleConnsKey, "10")
	t.Setenv(testMaxOpenConnsKey, "100")
}

// TestConfig_GetConfig tests the configuration loading
func TestConfig_GetConfig(t *testing.T) {
	setup(t)

	// Reset configInstance for a clean state
	configInstance = nil
	config, err := GetConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)

	assert.Equal(t, "9090", config.ServerPort)
	assert.Equal(t, "postgres://user:password@localhost:5432/mydb", config.DatabaseURI)
	assert.Equal(t, 10, config.MaxIdleConns)  // New field
	assert.Equal(t, 100, config.MaxOpenConns) // New field
}

// TestConfig_GetConfigWithoutDbURI tests configuration loading without database URI
func TestConfig_GetConfigWithoutDbURI(t *testing.T) {
	configInstance = nil
	os.Unsetenv(testDbURIKey) // Unset DB_URI

	_, err := GetConfig()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing required environment variable: DB_URI")
}

// TestConfig_GetConfigWithoutMaxIdleConns tests loading configuration without MAX_IDLE_CONNS
func TestConfig_GetConfigWithoutMaxIdleConns(t *testing.T) {
	setup(t)
	os.Unsetenv(testMaxIdleConnsKey) // Unset MAX_IDLE_CONNS

	configInstance = nil
	_, err := GetConfig()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing required environment variable: MAX_IDLE_CONNS")
}

// TestConfig_GetConfigWithoutMaxOpenConns tests loading configuration without MAX_OPEN_CONNS
func TestConfig_GetConfigWithoutMaxOpenConns(t *testing.T) {
	setup(t)
	configInstance = nil
	os.Unsetenv(testMaxOpenConnsKey) // Unset MAX_OPEN_CONNS

	_, err := GetConfig()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing required environment variable: MAX_OPEN_CONNS")
}

// TestConfig_GetConfigWithDefault tests loading configuration with default values
func TestConfig_GetConfigWithDefault(t *testing.T) {
	t.Setenv(testDbURIKey, "postgres://user:password@localhost:5432/mydb")
	t.Setenv(testMaxIdleConnsKey, "0")
	t.Setenv(testMaxOpenConnsKey, "0")

	// Do not set MAX_IDLE_CONNS or MAX_OPEN_CONNS to check for defaults
	t.Cleanup(func() {
		os.Unsetenv(testDbURIKey)
	})

	configInstance = nil
	config, err := GetConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)

	assert.Equal(t, "8080", config.ServerPort) // Default value
	assert.NotEmpty(t, config.DatabaseURI)     // Should not be empty as DB_URI is set
	assert.Equal(t, 0, config.MaxIdleConns)    // Default should be 0 since it was unset
	assert.Equal(t, 0, config.MaxOpenConns)    // Default should be 0 since it was unset
}
