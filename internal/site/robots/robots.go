package robots

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// Register installs the static robots.txt route on the global MVC router.
func Register() {
	mvc.StaticFile("/robots.txt", "robots/robots.txt")
}
