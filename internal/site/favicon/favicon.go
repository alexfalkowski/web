package favicon

import "github.com/alexfalkowski/web/internal/site/static"

// Register installs the static favicon route on the global MVC router.
func Register() {
	static.File("/favicon.ico", "favicon/favicon.png")
}
