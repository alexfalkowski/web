package static

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

const cacheControl = "public, max-age=3600"

// File registers an embedded static asset with the site's cache policy.
//
// Keep the shared MVC weak metadata validator intentionally. The embedded
// assets are stable between deploys, and this avoids repo-local hashing or
// buffering.
func File(pattern, name string) bool {
	return mvc.StaticFile(
		pattern,
		name,
		mvc.WithCacheControl(cacheControl),
		mvc.WithCacheValidators(),
	)
}
