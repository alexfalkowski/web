package favicon

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// Register installs the static favicon route on the global MVC router.
func Register() {
	mvc.StaticFile(
		"/favicon.ico",
		"favicon/favicon.png",
		mvc.WithCacheControl("public, max-age=3600"),
		mvc.WithCacheValidators(),
	)
}
