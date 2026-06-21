package sitemap

import "github.com/alexfalkowski/web/internal/site/static"

// Register installs the static sitemap.xml route on the global MVC router.
func Register() {
	static.File("/sitemap.xml", "sitemap/sitemap.xml")
}
