package meta

import (
	"time"

	"github.com/alexfalkowski/go-service/v2/env"
)

// NewYear returns the current calendar year.
//
// The value is derived from the local system clock at the time this function is
// called. It is typically constructed once during service startup and injected
// into view models via dependency injection.
func NewYear() Year {
	return Year(time.Now().Year())
}

// Year represents a calendar year (for example 2026).
//
// It is used as display metadata in templates (for example in a footer).
type Year int

// NewInfo constructs an Info value from the service version and year.
//
// The returned value is intended to be shared across multiple page models so
// templates can render consistent metadata.
func NewInfo(version env.Version, year Year) *Info {
	return &Info{Version: version, Year: year}
}

// Info contains metadata shared by multiple site view models.
type Info struct {
	// Version is the running service version (as provided by the environment/config).
	Version env.Version

	// Year is the current calendar year for display in templates.
	Year Year
}
