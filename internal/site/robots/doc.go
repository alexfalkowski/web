// Package robots registers the robots.txt endpoint for the site.
//
// The robots.txt file communicates crawler directives to search engines and other
// automated agents. This package exposes it as a static file at:
//
//   - GET /robots.txt
//
// The file content is bundled with the binary (via the site embedded filesystem),
// allowing the service to run without relying on external assets at runtime.
//
// # Public API
//
// Register installs the static route for /robots.txt on the global MVC router.
// It is typically invoked via dependency injection (see this package's Module)
// rather than being called directly by application code.
package robots
