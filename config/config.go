package config

import (
	"sync"

	envx "go.strv.io/env"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

const dotenvPath = ".env"

var (
	once sync.Once

	validate = validator.New()
)

type Config struct {
	Port         int    `env:"PORT" validate:"required"`
	DatabaseURL  string `env:"DATABASE_URL" validate:"required"`
	ResendApiKey string `env:"RESEND_API_KEY" validate:"required"`
	SupabaseURL        string `env:"SUPABASE_URL" validate:"required"`
	SupabaseAuthSecret string `env:"SUPABASE_AUTH_SECRET" validate:"required"`
	SupabaseAPIKey     string `env:"SUPABASE_API_KEY" validate:"required"`
}

func LoadConfig() (Config, error) {
	loaddotenv(dotenvPath)

	cfg := Config{}
	if err := envx.Apply(&cfg); err != nil {
		return cfg, err
	}
	if err := validate.Struct(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func MustLoadConfig() Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

func loaddotenv(path string) {
	once.Do(func() {
		if path == "" {
			path = ".env"
		}

		_ = godotenv.Load(dotenvPath)
		_ = godotenv.Load(dotenvPath + ".common")
		// Load local environment settings which may override previous settings
		_ = godotenv.Load(path + ".local") //
	})
}
