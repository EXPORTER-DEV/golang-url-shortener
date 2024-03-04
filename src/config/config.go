package config

import (
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	PRODUCTION_ENVIRONMENT  = "PRODUCTION"
	SANDBOX_ENVIRONMENT     = "SANDBOX"
	DEVELOPMENT_ENVIRONMENT = "DEVELOPMENT"
)

type Config struct {
	ENVIRONMENT    string `env:"ENVIRONMENT"`
	REDIS_ADDRESS  string `env:"REDIS_ADDR" env-required:"true"`
	REDIS_PASSWORD string `env:"REDIS_PASSWORD" env-required:"true"`
	REDIS_DATABASE int    `env:"REDIS_DATABASE" env-required:"true"`
	PORT           int    `env:"PORT"`
}

func (config *Config) IsProduction() bool {
	return config.ENVIRONMENT == PRODUCTION_ENVIRONMENT
}

func (config *Config) IsSandbox() bool {
	return config.ENVIRONMENT == SANDBOX_ENVIRONMENT
}

func New() (*Config, error) {
	config := Config{}

	if os.Getenv("ENVIRONMENT") == "" {
		wd, err := os.Getwd()

		if err != nil {
			return nil, err
		}

		var path string

		// Need following hack to run via VS Code profile, where working directory is src:
		if filepath.Base(wd) == "src" {
			path = filepath.Join(wd, "../.env")
		} else {
			path = filepath.Join(wd, ".env")
		}

		godotenv.Load(path)
	}

	err := cleanenv.ReadEnv(&config)

	if config.PORT == 0 {
		config.PORT = 80
	}

	return &config, err
}
