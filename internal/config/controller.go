package config

import (
	"github.com/caarlos0/env/v8"
	"go.uber.org/zap"
)

const (
	defaultBindAddr = "localhost:8080"
)

// ControllerConfig configures web server.
type ControllerConfig struct {
	// Addr specifies address on which server will be running at.
	Addr string `env:"BIND_ADDR" envDefault:"localhost:8080"`
}

// NewControllerConfig initializes controller config and returns it to user.
func NewControllerConfig() (*ControllerConfig, error) {
	cfg := new(ControllerConfig)
	_ = env.Parse(cfg)
	return cfg, nil
}

// BindAddr returns address on which server will be running at.
func (cfg *ControllerConfig) BindAddr() string {
	if cfg == nil {
		zap.L().Warn("unexpectedly got nil pointer receiver config")
		return defaultBindAddr
	}
	return cfg.Addr
}
