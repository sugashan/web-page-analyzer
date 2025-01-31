package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test loading config from a valid JSON file
func TestLoadConfig_ValidFile(t *testing.T) {
	// Create a temporary config file
	tempFile := "test_config.json"
	configData := `{"request_timeout": 15}`
	err := os.WriteFile(tempFile, []byte(configData), 0644)
	assert.NoError(t, err)

	// Ensure the file is deleted after the test
	defer os.Remove(tempFile)

	// Load the config
	err = LoadConfig(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, 15, AppConfig.RequestTimeout)
	assert.Equal(t, 15*time.Second, GetRequestTimeout())
}

// Test loading config from a missing file
func TestLoadConfig_FileNotFound(t *testing.T) {
	err := LoadConfig("nonexistent_config.json")
	assert.Error(t, err)
}

// Test loading config with an invalid JSON format
func TestLoadConfig_InvalidJSON(t *testing.T) {
	// Create a temporary invalid config file
	tempFile := "invalid_config.json"
	invalidData := `{"request_timeout": "abc"}`
	err := os.WriteFile(tempFile, []byte(invalidData), 0644)
	assert.NoError(t, err)

	// Ensure the file is deleted after the test
	defer os.Remove(tempFile)

	// Load the config
	err = LoadConfig(tempFile)
	assert.Error(t, err)
}
