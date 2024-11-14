package robots

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register robots.
func Register(router *mvc.Router) {
	router.Static("GET /robots.txt", "robots/robots.txt")
}
