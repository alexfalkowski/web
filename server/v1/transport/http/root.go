package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// RootModel for HTTP.
type RootModel struct {
	Title string
}

// RootRoute registers routes for root (index).
func RootRoute() {
	v := mvc.NewView(mvc.ParseTemplate(Views, "root.tmpl.html"), nil)

	mvc.Route("GET /", v, func(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*RootModel, error) {
		m := &RootModel{Title: "Lean Thoughts"}

		return m, nil
	})
}
