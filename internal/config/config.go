package config

import (
	"github.com/alexfalkowski/go-service/v2/config"
	"github.com/alexfalkowski/web/internal/health"
)

// Config is the root configuration for the web service.
//
// It embeds the shared base configuration (`go-service/v2/config.Config`) and adds
// service-specific sections (for example health probe configuration).
//
// The struct tags allow it to be loaded from YAML, JSON, or TOML.
type Config struct {
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}
