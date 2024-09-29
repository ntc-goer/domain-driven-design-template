package configs

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port int `env:"PORT, default=8080"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Load(ctx context.Context) error {
	if err := envconfig.Process(ctx, c); err != nil {
		return err
	}
	return nil
}
