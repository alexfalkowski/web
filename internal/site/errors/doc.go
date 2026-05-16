// Package errors wires the site's error feature.
//
// The error feature provides the MVC not-found renderer used when no route
// matches a request. It is composed into the wider site using dependency
// injection.
//
// Consumers typically do not call into this package directly. Instead, they
// include `site.Module`, which includes `errors.Module` and thereby installs the
// not-found registration.
//
// # Public API
//
// The public API of this package is `Module`, a DI module that installs the
// error feature into the application.
package errors
