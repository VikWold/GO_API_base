package config

import (
	"errors"
	"os"
)

type Config struct {
	DB *DatabaseConfig
}

type DatabaseConfig struct {
	DSN     string `json:"-"`
	Timeout int    `json:"timeout"`
}

func New() (*Config, error) {
	dsn := os.Getenv("CHANGE_NAME_DB_DSN")
	if dsn == "" {
		return nil, errors.New("CHANGE_NAME_DB_DSN environment variable is not set")
	}

	return &Config{
		DB: &DatabaseConfig{
			DSN:     dsn,
			Timeout: 30,
		},
	}, nil
}
