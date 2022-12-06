package config

import "os"

type Config struct {
	CredentialsBin []byte
}

func New() (*Config, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, nil
	}
	return &Config{
		CredentialsBin: b,
	}, nil
}
