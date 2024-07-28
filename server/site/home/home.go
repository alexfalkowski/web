package home

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register home.
func Register(router *mvc.Router) {
	router.Route("GET /home", func(_ context.Context) (mvc.View, mvc.Model) {
		return mvc.View("home.tmpl"), nil
	})
}
