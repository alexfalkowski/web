// Package sitemap registers the site's sitemap.xml endpoint.
//
// The sitemap.xml file lists the site's canonical public URLs for search
// engines and other crawlers. This package exposes it as a static file at:
//
//   - GET /sitemap.xml
//
// The file content is bundled with the binary through the site embedded
// filesystem, allowing the service to run without relying on external assets at
// runtime.
//
// # Public API
//
// Register installs the static route for /sitemap.xml on the global MVC router.
// It is typically invoked via dependency injection (see this package's Module)
// rather than being called directly by application code.
package sitemap
