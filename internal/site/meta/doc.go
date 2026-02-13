// Package meta provides shared, cross-cutting metadata used by site view models.
//
// The types in this package are used to enrich models with information that is
// common across pages, such as the running service version and the current year
// for rendering in templates (for example in footers or page headers).
//
// Values are typically constructed via dependency injection:
//
//   - NewYear produces a Year based on the current system time.
//   - NewInfo assembles an Info struct from the injected service version and year.
//
// # Public API
//
//   - Year represents the current calendar year.
//   - Info contains metadata shared by multiple site models (Version and Year).
//   - NewYear and NewInfo are constructors intended to be used by the DI graph.
//
// This package does not perform any I/O beyond reading the current time.
package meta
