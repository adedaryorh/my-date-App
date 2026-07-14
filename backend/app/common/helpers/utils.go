// GetenvAsInt gets an environment variable as an integer, returns defaultValue if not found or invalid
package helpers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// GetenvAsInt gets an environment variable as an integer, returns defaultValue if not found or invalid
func GetenvAsInt(key string, defaultValue int) int {
	valueStr := Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

type Map map[string]interface{}

const randomAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GetDurationFromTimeString(value string) time.Duration {
	duration, _ := time.ParseDuration(value)
	return duration
}
func GenerateRandomNumber(length int) string    { return generateRandom(length, 10) }
func GenerateRandomByte(length int) string      { return generateRandom(length, len(randomAlphabet)) }
func GenerateRandomUppercase(length int) string { return strings.ToUpper(GenerateRandomByte(length)) }
func generateRandom(length, limit int) string {
	result := make([]byte, length)
	for index := range result {
		result[index] = randomAlphabet[rand.Intn(limit)]
	}
	return string(result)
}
func HashString(value string) string {
	sum := sha256.Sum256([]byte(value))
	return base64.StdEncoding.EncodeToString(sum[:])
}
func Hash(value string) string {
	encoded, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(encoded)
}
func CompareHash(hashed, value string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(value)) == nil
}
func PointerString(value string) *string { return &value }
func GetEntityComputedHash(entity interface{}) (map[string]interface{}, string) {
	data, _ := json.Marshal(entity)
	values := map[string]interface{}{}
	_ = json.Unmarshal(data, &values)
	return values, HashString(string(data))
}
func GetFileType(filename, separator string) string {
	parts := strings.Split(filename, separator)
	if len(parts) == 0 {
		return ""
	}
	return strings.ToLower(parts[len(parts)-1])
}

func GetenvAsFloat(key string, defaultValue float64) float64 {
	if value, err := strconv.ParseFloat(Getenv(key), 64); err == nil {
		return value
	}
	return defaultValue
}

func Int64ToPointer(value int64) *int64 { return &value }

func Reverse[T any](values []T) []T {
	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		values[left], values[right] = values[right], values[left]
	}
	return values
}

func ParseTemplate(path string, data interface{}) (*string, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}
	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		return nil, err
	}
	result := output.String()
	return &result, nil
}

// Getenv gets an environment variable or returns an empty string if not found
func Getenv(key string, defaults ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return ""
}
