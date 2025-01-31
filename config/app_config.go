package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	RequestTimeout int `json:"request_timeout"` // Timeout in seconds
}

var AppConfig Config

// LoadConfig reads configuration from a JSON file
func LoadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}

// GetRequestTimeout returns the timeout as a time.Duration
func GetRequestTimeout() time.Duration {
	return time.Duration(AppConfig.RequestTimeout) * time.Second
}
