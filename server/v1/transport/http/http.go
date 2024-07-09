package http

import (
	"context"
	"embed"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

//go:embed views/index.tmpl.html
var views embed.FS

// Register for http.
func Register() {
	v := mvc.NewView(mvc.ParseTemplate(views, "views/index.tmpl.html"), nil)

	mvc.Route("GET /", v, func(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*string, error) {
		s := "Hello"

		return &s, nil
	})
}
