package sitemap

import (
	"github.com/alexfalkowski/go-service/v3/net/http/mvc"
	"github.com/alexfalkowski/web/v2/internal/site/static"
)

// Register installs the static sitemap.xml route on the global MVC router.
func Register(server *mvc.Server) {
	static.File(server, "/sitemap.xml", "sitemap/sitemap.xml")
}
