package v1

import (
	"context"
	"embed"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

//go:embed *.html
var fs embed.FS

// Register for v1.
func Register() {
	mvc.Route(
		"GET /",
		mvc.NewSuccessView(mvc.ParseTemplate(fs, "root.html")),
		mvc.NoController,
	)

	mvc.Route(
		"GET /home",
		mvc.NewSuccessView(mvc.ParseTemplate(fs, "home.html")),
		mvc.NoController,
	)

	mvc.Route(
		"GET /books",
		mvc.NewSuccessView(mvc.ParseTemplate(fs, "books.html")),
		func(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*BookModel, error) {
			m := &BookModel{Books: books()}

			return m, nil
		},
	)
}
