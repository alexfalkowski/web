package home

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Model for home.
type Model struct{}

// Register home.
func Register() {
	mvc.Route("GET /home", func(_ context.Context) (mvc.View, *Model, error) {
		m := &Model{}

		return mvc.View("home.tmpl"), m, nil
	})
}
