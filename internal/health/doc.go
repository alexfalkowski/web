// Package health configures and exposes health and liveness/readiness endpoints.
//
// This package registers health checks and HTTP observers against the service's
// server, enabling standard Kubernetes-style probes:
//
//   - /healthz   overall health (online)
//   - /livez     liveness (noop)
//   - /readyz    readiness (noop)
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
