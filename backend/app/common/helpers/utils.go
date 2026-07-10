// GetenvAsInt gets an environment variable as an integer, returns defaultValue if not found or invalid
func GetenvAsInt(key string, defaultValue int) int {
	valueStr := Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}