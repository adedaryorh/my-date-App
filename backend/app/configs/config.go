package configs

import (
	"log"
	"os"
	"strconv"

	"backend.app/common/helpers"
	validator "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string `validate:"required" yaml:"name"    env:"APP_NAME"`
	LogLevel   string `validate:"required" yaml:"log_level"   env:"LOG_LEVEL"`
	AppVersion string `validate:"required" yaml:"version" env:"APP_VERSION"`
	AppEnv     string `validate:"required" yaml:"app_env" env:"APP_ENV"`
	AppHost    string `validate:"required" yaml:"app_host" env:"APP_HOST"`
	Port       string `validate:"required" yaml:"port" env:"PORT"`
	GinMode    string

	PGHost     string `validate:"required"`
	PGPort     string `validate:"required"`
	PGUser     string `validate:"required"`
	PGPassword string `validate:"required"`
	PGDatabase string `validate:"required"`
	PGSSlMode  string
	PgPoolMax  string `validate:"required" yaml:"pg_pool_max" env:"PG_POOL_MAX"`

	RedisUri string `validate:"required" yaml:"redis_uri" env:"REDIS_URI"`

	JwtSecret       string `validate:"required" yaml:"jwt_secret" env:"JWT_SECRET"`
	JwtSecretExpiry string `validate:"required" yaml:"jwt_secret_expiry" env:"JWT_SECRET_EXPIRY"`

	ServiceAddress string `validate:"required" yaml:"service_address" env:"SERVICE_ADDRESS"`

	AwsAccessKeyID     string `validate:"required" yaml:"access_key_id" env:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `validate:"required" yaml:"secret_id" env:"AWS_SECRET_ID"`
	AwsRegion          string `validate:"required" yaml:"region" env:"AWS_REGION"`
	AwsS3Bucket        string `validate:"required" yaml:"aws_s3_bucket" env:"AWS_S3_BUCKET"`

	EnableSwagger            string `validate:"required"`
	SendGridApiKey           string `validate:"required"`
	TwilioAccountSid         string `validate:"required"`
	TwilioAuthToken          string `validate:"required"`
	TwilioMessagingServiceId string `validate:"required"`
	MixPanelProjectToken     string `validate:"required"`
	MixPanelApiSecret        string `validate:"required"`
	MixPanelProjectId        int    `validate:"required"`
	MixPanelUsername         string `validate:"required"`

	AIServiceURL                     string `validate:"required" yaml:"ai_service_url" env:"AI_SERVICE_URL"`
	AIServiceTimeout                 string `validate:"required" yaml:"ai_service_timeout" env:"AI_SERVICE_TIMEOUT"` // e.g., "5s"
	AIServiceMaxRetries              int    `validate:"required" yaml:"ai_service_max_retries" env:"AI_SERVICE_MAX_RETRIES"`
	AIServiceCircuitBreakerThreshold int    `validate:"validate:"required" yaml:"ai_service_circuit_breaker_threshold" env:"AI_SERVICE_CIRCUIT_BREAKER_THRESHOLD"` // Number of consecutive failures before opening circuit

	// Additional circuit breaker settings
	AIServiceCircuitBreakerTimeout   int    `validate:"required" yaml:"ai_service_circuit_breaker_timeout" env:"AI_SERVICE_CIRCUIT_BREAKER_TIMEOUT"` // Seconds before trying half-open state
	AIServiceCircuitBreakerInterval  int    `validate:"required" yaml:"ai_service_circuit_breaker_interval" env:"AI_SERVICE_CIRCUIT_BREAKER_INTERVAL"` // Interval between state changes
	AIServiceCircuitBreakerMaxRequests int   `validate:"required" yaml:"ai_service_circuit_breaker_max_requests" env:"AI_SERVICE_CIRCUIT_BREAKER_MAX_REQUESTS"` // Max requests in half-open state
	AIServiceFailureRateThreshold    float64 `validate:"required" yaml:"ai_service_failure_rate_threshold" env:"AI_SERVICE_FAILURE_RATE_THRESHOLD"` // Failure rate to trip circuit (0.0-1.0)

	GoogleClientID     string `validate:"required" yaml:"google_client_id" env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `validate:"required" yaml:"google_client_secret" env:"GOOGLE_CLIENT_SECRET"`
}

func NewConfig() (*Config, error) {

	if os.Getenv("APP_ENV") != "prod" && os.Getenv("APP_ENV") != "stg" && os.Getenv("APP_ENV") != "beta" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("env file error: %s", err.Error())
		}
	}
	mixPanelProjectId, err := strconv.Atoi(helpers.Getenv("MIX_PANEL_PROJECT_ID"))
	if err != nil {
		log.Fatalf("mix panel project id error: %s", err.Error())
	}
	config := Config{
		AppName:    helpers.Getenv("APP_NAME"),
		AppEnv:     helpers.Getenv("APP_ENV", "local"),
		LogLevel:   helpers.Getenv("LOG_LEVEL", "debug"),
		AppVersion: helpers.Getenv("APP_VERSION", "1.0"),
		AppHost:    helpers.Getenv("APP_HOST", "0.0.0.0"),
		Port:       helpers.Getenv("PORT", "7070"),
		RedisUri:   helpers.Getenv("REDIS_URI"),
		PGHost:     os.Getenv("PG_HOST"),
		PGPort:     os.Getenv("PG_PORT"),
		PGUser:     os.Getenv("PG_USER"),
		PGPassword: os.Getenv("PG_PASSWORD"),
		PGDatabase: os.Getenv("PG_DATABASE"),
		//RedisPort:               helpers.Getenv("REDIS_PORT"),
		EnableSwagger:            helpers.Getenv("ENABLE_SWAGGER", "true"),
		AwsRegion:                helpers.Getenv("AWS_REGION"),
		AwsAccessKeyID:           helpers.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey:       helpers.Getenv("AWS_SECRET_ACCESS_KEY"),
		AwsS3Bucket:              helpers.Getenv("AWS_S3_BUCKET"),
		ServiceAddress:           helpers.Getenv("SERVICE_ADDRESS"),
		PgPoolMax:                helpers.Getenv("PG_POOL_MAX", "2"),
		JwtSecret:                helpers.Getenv("JWT_SECRET"),
		JwtSecretExpiry:          helpers.Getenv("JWT_SECRET_EXPIRY"),
		GinMode:                  helpers.Getenv("GIN_MODE"),
		SendGridApiKey:           helpers.Getenv("SEND_GRID_API_KEY"),
		TwilioAccountSid:         helpers.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:          helpers.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioMessagingServiceId: helpers.Getenv("TWILIO_MESSAGING_SERVICE_ID"),
		MixPanelProjectToken:     helpers.Getenv("MIX_PANEL_PROJECT_TOKEN"),
		MixPanelApiSecret:        helpers.Getenv("MIX_PANEL_API_SECRET"),
		MixPanelUsername:         helpers.Getenv("MIX_PANEL_USERNAME"),
		MixPanelProjectId:        mixPanelProjectId,
		// AI Service Configuration
		AIServiceURL:                     helpers.Getenv("AI_SERVICE_URL", "http://localhost:8000"),
		AIServiceTimeout:                 helpers.Getenv("AI_SERVICE_TIMEOUT", "5s"),
		AIServiceMaxRetries:              helpers.GetenvAsInt("AI_SERVICE_MAX_RETRIES", 3),
		AIServiceCircuitBreakerThreshold: helpers.GetenvAsInt("AI_SERVICE_CIRCUIT_BREAKER_THRESHOLD", 5),

		// Additional circuit breaker settings with defaults
		AIServiceCircuitBreakerTimeout:   helpers.GetenvAsInt("AI_SERVICE_CIRCUIT_BREAKER_TIMEOUT", 60),
		AIServiceCircuitBreakerInterval:  helpers.GetenvAsInt("AI_SERVICE_CIRCUIT_BREAKER_INTERVAL", 10),
		AIServiceCircuitBreakerMaxRequests: helpers.GetenvAsInt("AI_SERVICE_CIRCUIT_BREAKER_MAX_REQUESTS", 3),
		AIServiceFailureRateThreshold:    helpers.GetenvAsFloat("AI_SERVICE_FAILURE_RATE_THRESHOLD", 0.5),
		GoogleClientID:     helpers.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: helpers.Getenv("GOOGLE_CLIENT_SECRET"),
	}

	validate := validator.New()
	if err = validate.Struct(config); err != nil {
		log.Fatalf("env validation error: %s", err.Error())
	}

	return &config, nil
}

// Helper function to get float64 from env
func GetenvAsFloat(key string, defaultValue float64) float64 {
	if value := os.Getenv(key); value != "" {
		if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
			return floatVal
		}
	}
	return defaultValue
}