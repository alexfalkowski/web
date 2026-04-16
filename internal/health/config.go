package health

import "github.com/alexfalkowski/go-service/v2/time"

// Config defines health check timing configuration.
//
// Duration and Timeout are duration strings (for example "5s", "1m") that are
// parsed as Go durations.
//
// Duration controls how often health registrations are evaluated/updated.
// Timeout is reserved for probe/check timeout configuration.
type Config struct {
	// Duration is the health check evaluation interval (for example "5s").
	Duration time.Duration `yaml:"duration,omitempty" json:"duration,omitempty" toml:"duration,omitempty"`

	// Timeout is the maximum time allowed for a health check/probe (for example "2s").
	Timeout time.Duration `yaml:"timeout,omitempty" json:"timeout,omitempty" toml:"timeout,omitempty"`
}
