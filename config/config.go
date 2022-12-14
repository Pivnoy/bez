package config

import (
	"github.com/caarlos0/env/v6"
	"os"
)

type Config struct {
	AppPort        string `env:"APP_PORT" envDefault:"9000"`
	PostgresUrl    string `env:"POSTGRES_URL" envDefault:"postgresql://postgres:postgres@localhost:5432/postgres"`
	CredentialsBin []byte
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile("./config/credentials.json")
	if err != nil {
		return nil, nil
	}
	cfg.CredentialsBin = b
	return cfg, nil
}
