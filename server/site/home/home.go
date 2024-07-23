package home

import (
	"context"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register home.
func Register(fs fs.FS) {
	mvc.Route("GET /home", func(_ context.Context, _ *http.Request, _ http.ResponseWriter) *mvc.Result {
		v := template.Must(template.ParseFS(fs, "home/view.html"))
		r := mvc.NewResult(nil, v)

		return r
	})
}
