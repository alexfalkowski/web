package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// HomeModel for HTTP.
type HomeModel struct {
	Title string
}

// HomeRoute registers routes for home.
func HomeRoute() {
	v := mvc.NewView(mvc.ParseTemplate(Views, "home.tmpl.html"), nil)

	mvc.Route("GET /", v, func(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*HomeModel, error) {
		h := &HomeModel{Title: "Lean Thoughts"}

		return h, nil
	})
}
