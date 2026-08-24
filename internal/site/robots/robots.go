package robots

import (
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/static"
)

// Register installs the static robots.txt route on the global MVC router.
func Register(server *mvc.Server) {
	static.File(server, "/robots.txt", "robots/robots.txt")
}
