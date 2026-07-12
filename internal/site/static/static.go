package static

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

const cacheControl = "public, max-age=3600"

// File registers a GET route for an embedded static asset. Responses use
// Cache-Control: public, max-age=3600 and weak metadata validators; matching
// conditional requests can return 304 Not Modified.
//
// File returns false when MVC has not been initialized and true after the route
// is registered.
//
// Keep the shared MVC weak metadata validator intentionally. The embedded
// assets are stable for the lifetime of a running binary, and this avoids
// repo-local hashing or buffering.
func File(pattern, name string) bool {
	return mvc.StaticFile(
		pattern,
		name,
		mvc.WithCacheControl(cacheControl),
		mvc.WithCacheValidators(),
	)
}
