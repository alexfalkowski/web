// Package route registers the global not-found handler for the error feature.
//
// This package installs the MVC not-found controller used when no registered
// route matches a request. The controller renders either the full not-found page
// or the partial not-found fragment, matching the same full/partial rendering
// behavior used by the normal site routes.
//
// # Public API
//
// Register installs the global not-found handler on the MVC router. It is
// typically invoked via dependency injection (see this package's Module) rather
// than being called directly by application code.
package route
