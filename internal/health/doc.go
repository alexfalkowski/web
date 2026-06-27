// Package health configures and exposes health and liveness/readiness endpoints.
//
// This package registers health checks and HTTP observers against the service's
// server. The HTTP transport exposes those observers on service-prefixed
// operation routes; for this service's local name, the routes are:
//
//   - /web/healthz   overall health (default online connectivity check)
//   - /web/livez     liveness (noop)
//   - /web/readyz    readiness (noop)
//
// The /web/healthz observer is backed by go-health/v2's default online
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
