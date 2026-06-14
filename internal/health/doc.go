// Package health configures and exposes health and liveness/readiness endpoints.
//
// This package registers health checks and HTTP observers against the service's
// server, enabling standard Kubernetes-style probes:
//
//   - /healthz   overall health (default online connectivity check)
//   - /livez     liveness (noop)
//   - /readyz    readiness (noop)
//
// The /healthz observer is backed by go-health/v2's default online
// registration, so restricted public connectivity can affect the overall health
// response. The liveness and readiness observers use noop registrations.
//
// The configuration for probe timing is provided via Config and is expected to
// be loaded from the service's root configuration.
//
// # Public API
//
// The public API of this package is:
//
//   - Config: YAML/JSON/TOML configuration for health check intervals/timeouts.
//   - Module: a DI module that registers checks and HTTP observers.
package health
