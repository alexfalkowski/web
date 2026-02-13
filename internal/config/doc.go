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
// The only public type in this package is Config, which represents the root of
// the service configuration file.
package config
