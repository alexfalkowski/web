package favicon

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// Register favicon.
func Register() {
	mvc.StaticFile("/favicon.ico", "favicon/favicon.png")
}
