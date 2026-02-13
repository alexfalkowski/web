// Package root wires the site's root (home) feature.
//
// The root feature provides the HTTP routes, controllers, views, and models used
// to render the site's home page. It is composed into the wider site using
// dependency injection.
//
// Consumers typically do not call into this package directly. Instead, they
// include `site.Module`, which includes `root.Module` and thereby installs the
// root route registration.
//
// # Public API
//
// The public API of this package is `Module`, a DI module that installs the
// root feature into the application.
package root
