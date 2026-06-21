package robots

import "github.com/alexfalkowski/web/internal/site/static"

// Register installs the static robots.txt route on the global MVC router.
func Register() {
	static.File("/robots.txt", "robots/robots.txt")
}
