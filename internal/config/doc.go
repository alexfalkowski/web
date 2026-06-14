// Package config defines the configuration model for the web service.
//
// The config package provides the service-specific configuration struct that is
// loaded from external configuration sources (for example YAML/JSON/TOML) and
// then used to construct the dependency graph.
//
// It embeds the shared base configuration from `go-service` and adds additional
// sections that are specific to this service (such as health checking).
//
// # Public API
//
// The public surface is:
//   - Config, which represents the root of the service configuration file.
//   - Module, which loads Config through the shared go-service configuration
//     constructor and exposes the embedded base config plus health config to the
//     DI graph.
package config
