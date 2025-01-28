package configs

import (
	"log"
	"os"

	"backend.app/common/helpers"
	validator "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		AppName    string `validate:"required" yaml:"name"    env:"APP_NAME"`
		LogLevel   string `validate:"required" yaml:"log_level"   env:"LOG_LEVEL"`
		AppVersion string `validate:"required" yaml:"version" env:"APP_VERSION"`
		AppEnv     string `validate:"required" yaml:"app_env" env:"APP_ENV"`
		AppHost    string `validate:"required" yaml:"app_host" env:"APP_HOST"`
		Port       string `validate:"required" yaml:"port" env:"PORT"`
		GinMode    string

		// PGHost     string `validate:"required"`
		// PGPort     string `validate:"required"`
		// PGUser     string `validate:"required"`
		// PGPassword string `validate:"required"`
		// PGDatabase string `validate:"required"`
		PgPoolMax string `validate:"required" yaml:"pg_pool_max" env:"PG_POOL_MAX"`
		PgUrl     string `validate:"required" yaml:"url"      env:"PG_URL"`

		RedisUri string `validate:"required" yaml:"redis_uri" env:"REDIS_URI"`

		JwtSecret       string `validate:"required" yaml:"jwt_secret" env:"JWT_SECRET"`
		JwtSecretExpiry string `validate:"required" yaml:"jwt_secret_expiry" env:"JWT_SECRET_EXPIRY"`

		ServiceAddress string `validate:"required" yaml:"service_address" env:"SERVICE_ADDRESS"`

		AwsAccessKeyID     string `validate:"required" yaml:"access_key_id" env:"AWS_ACCESS_KEY_ID"`
		AwsSecretAccessKey string `validate:"required" yaml:"secret_id" env:"AWS_SECRET_ID"`
		AwsRegion          string `validate:"required" yaml:"region" env:"AWS_REGION"`
		StorageEndpoint    string
		AwsS3Bucket        string `validate:"required" yaml:"aws_s3_bucket" env:"AWS_S3_BUCKET"`

		EnableSwagger string `validate:"required"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {

	if os.Getenv("APP_ENV") != "prod" && os.Getenv("APP_ENV") != "stg" && os.Getenv("APP_ENV") != "beta" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("env file error: %s", err.Error())
		}
	}

	config := Config{
		AppName:    helpers.Getenv("APP_NAME"),
		AppEnv:     helpers.Getenv("APP_ENV", "local"),
		LogLevel:   helpers.Getenv("LOG_LEVEL", "debug"),
		AppVersion: helpers.Getenv("APP_VERSION", "1.0"),
		AppHost:    helpers.Getenv("APP_HOST", "0.0.0.0"),
		Port:       helpers.Getenv("PORT", "7070"),
		RedisUri:   helpers.Getenv("REDIS_URI"),
		//RedisPort:               helpers.Getenv("REDIS_PORT"),
		EnableSwagger:      helpers.Getenv("ENABLE_SWAGGER", "true"),
		AwsRegion:          helpers.Getenv("AWS_REGION"),
		AwsAccessKeyID:     helpers.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: helpers.Getenv("AWS_SECRET_ACCESS_KEY"),
		AwsS3Bucket:        helpers.Getenv("AWS_S3_BUCKET"),
		StorageEndpoint:    helpers.Getenv("STORAGE_ENDPOINT"),
		ServiceAddress:     helpers.Getenv("SERVICE_ADDRESS"),
		PgPoolMax:          helpers.Getenv("PG_POOL_MAX", "2"),
		PgUrl:              helpers.Getenv("PG_URL"),
		JwtSecret:          helpers.Getenv("JWT_SECRET"),
		JwtSecretExpiry:    helpers.Getenv("JWT_SECRET_EXPIRY"),
		GinMode:            helpers.Getenv("GIN_MODE"),
	}

	validate := validator.New()
	err := validate.Struct(config)

	if err != nil {
		log.Fatalf("env validation error: %s", err.Error())
	}

	return &config, nil
}
