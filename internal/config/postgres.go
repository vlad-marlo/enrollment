package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type PgConfig struct {
	USER string `env:"POSTGRES_USER" envDefault:"postgres"`
	PASS string `env:"POSTGRES_PASS" envDefault:"postgres"`
	PORT int    `env:"POSTGRES_PORT" envDefault:"5432"`
	HOST string `env:"POSTGRES_HOST" envDefault:"localhost"`
	NAME string `env:"POSTGRES_NAME" envDefault:"postgres"`
}

const (
	defaultPGValue = "postgres"
	defaultPGAddr  = "localhost"
	defaultPGPort  = 5432
)

func NewPgConfig() (*PgConfig, error) {
	cfg := new(PgConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("error while parsing pg config: %w", err)
	}
	return cfg, nil
}

func (cfg *PgConfig) User() string {
	if cfg == nil {
		return defaultPGValue
	}
	return cfg.USER
}

func (cfg *PgConfig) Pass() string {
	if cfg == nil {
		return defaultPGValue
	}
	return cfg.PASS
}

func (cfg *PgConfig) Port() int {
	if cfg == nil {
		return defaultPGPort
	}
	return cfg.PORT
}

func (cfg *PgConfig) Host() string {
	if cfg == nil {
		return defaultPGAddr
	}
	return cfg.HOST
}

func (cfg *PgConfig) Name() string {
	if cfg == nil {
		return defaultPGValue
	}
	return cfg.NAME
}

func (cfg *PgConfig) URI() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", cfg.User(), cfg.Pass(), cfg.Host(), cfg.Port(), cfg.Name())
}
