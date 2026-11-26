package static

import "github.com/alexfalkowski/go-service/v3/net/http/mvc"

const cacheControl = "public, max-age=3600"

// File registers a GET route for an embedded static asset using the shared MVC
// handler and the site's public, one-hour cache policy.
//
// File returns false when MVC has not been initialized and true after the route
// is registered.
func File(server *mvc.Server, pattern, name string) bool {
	return server.StaticFile(
		pattern,
		name,
		mvc.WithCacheControl(cacheControl),
	)
}
