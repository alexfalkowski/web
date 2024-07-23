package root

import (
	"context"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register root.
func Register(fs fs.FS) {
	mvc.Route("GET /", func(_ context.Context, _ *http.Request, _ http.ResponseWriter) *mvc.Result {
		v := template.Must(template.ParseFS(fs, "root/view.html"))
		r := mvc.NewResult(nil, v)

		return r
	})
}
