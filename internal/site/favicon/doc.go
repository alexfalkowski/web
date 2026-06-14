// Package favicon registers the site's browser icon endpoint.
//
// The favicon image is bundled with the binary through the site embedded
// filesystem and exposed as a static file at:
//
//   - GET /favicon.ico
//
// # Public API
//
// Register installs the static route for /favicon.ico on the global MVC router.
// It is typically invoked via dependency injection (see this package's Module)
// rather than being called directly by application code.
package favicon
