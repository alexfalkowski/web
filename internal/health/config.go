package health

import "github.com/alexfalkowski/go-service/v2/time"

// Config defines health check timing configuration.
//
// Duration and Timeout are duration strings (for example "5s", "1m") that are
// parsed as Go durations.
//
// Duration controls how often health registrations are evaluated/updated.
// Timeout controls the default online health check timeout. A zero value selects
// the go-health/v2 default of 30 seconds.
type Config struct {
	// Duration is the health check evaluation interval and must be greater than
	// zero (for example "5s").
	Duration time.Duration `yaml:"duration,omitempty" json:"duration,omitempty" toml:"duration,omitempty" validate:"gt=0"`

	// Timeout is the maximum time allowed for the online health check and may be
	// zero or greater (for example "2s"). Zero selects the go-health/v2 default
	// of 30 seconds.
	Timeout time.Duration `yaml:"timeout,omitempty" json:"timeout,omitempty" toml:"timeout,omitempty" validate:"gte=0"`
}
