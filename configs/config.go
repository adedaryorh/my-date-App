package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		PG    `yaml:"postgres"`
		Token `yaml:"token"`
		OTP   `yaml:"otp"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Env     string `env-required:"true" yaml:"version" env:"ENV"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		ServiceAddress string `env-required:"false" yaml:"service_address" env:"SERVICE_ADDR"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"url"      env:"PG_URL"`
	}

	// Token -.
	Token struct {
		Secret string `env-required:"true" yaml:"token_secret" env:"TOKEN_SECRET"`
	}

	// OTP -.
	OTP struct {
		Secret string `env-required:"true" yaml:"otp_secret" env:"OTP_SECRET"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	//configPath := os.Getenv("CONFIG_ENV")
	//err := cleanenv.ReadConfig(configPath, cfg)

	//if err != nil {
	//	return nil, fmt.Errorf("config error: %w", err)
	//}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
