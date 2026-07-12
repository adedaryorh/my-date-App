// GetenvAsInt gets an environment variable as an integer, returns defaultValue if not found or invalid
package helpers

import (
	"os"
	"strconv"
)

// GetenvAsInt gets an environment variable as an integer, returns defaultValue if not found or invalid
func GetenvAsInt(key string, defaultValue int) int {
	valueStr := Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// Getenv gets an environment variable or returns an empty string if not found
func Getenv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}

// Getenv gets an environment variable or returns an empty string if not found
func Getenv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}