// Package route registers HTTP routes for the root (home) feature.
//
// This package wires the home page endpoints into the MVC router. Route handlers
// are implemented as controllers, which render views backed by the root model.
//
// Routes registered by this package:
//
//   - GET / renders the full home page.
//   - PUT / renders the partial home page (for incremental/HTMX-style updates),
//     using a partial layout/view pair.
//
// # Public API
//
// Register installs the root routes on the global MVC router. It is typically
// invoked via dependency injection (see this package's Module) rather than being
// called directly by application code.
package route
