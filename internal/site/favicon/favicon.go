package favicon

import (
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/static"
)

// Register installs the static favicon route on the global MVC router.
func Register(server *mvc.Server) {
	static.File(server, "/favicon.ico", "favicon/favicon.png")
}
