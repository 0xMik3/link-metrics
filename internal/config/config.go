package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port       string `env:"PORT" envDefault:"8080"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"debug"`
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbUser     string `env:"DB_USER" envDefault:"postgres"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DbName     string `env:"DB_NAME" envDefault:"postgres"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbSslMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

func (c *Config) GetEnvs(ctx context.Context) context.Context {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}
	envconfig.Process(ctx, c)
	ctx = context.WithValue(ctx, "envs", c)
	return ctx
}
