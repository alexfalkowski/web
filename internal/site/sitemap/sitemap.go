package sitemap

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// Register installs the static sitemap.xml route on the global MVC router.
func Register() {
	mvc.StaticFile(
		"/sitemap.xml",
		"sitemap/sitemap.xml",
		mvc.WithCacheControl("public, max-age=3600"),
		mvc.WithCacheValidators(),
	)
}
