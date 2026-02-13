// Package route registers HTTP routes for the books feature.
//
// This package wires the books endpoints into the MVC router. Route handlers are
// implemented as controllers, which render views backed by the books model.
//
// Routes registered by this package:
//
//   - GET /books renders the full books page.
//   - PUT /books renders the partial books page (for incremental/HTMX-style
//     updates), using a partial layout/view pair.
//
// # Public API
//
// Register installs the books routes on the global MVC router. It is typically
// invoked via dependency injection (see this package's Module) rather than being
// called directly by application code.
package route
