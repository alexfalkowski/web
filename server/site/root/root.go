package root

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register root.
func Register(router *mvc.Router) {
	router.Route("GET /", func(_ context.Context) (mvc.View, mvc.Model) {
		return mvc.View("root.tmpl"), nil
	})
}
