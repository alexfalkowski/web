package root

import (
	"context"
	"html/template"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register root.
func Register(fs fs.FS) {
	mvc.Route("GET /", func(_ context.Context) (*template.Template, any) {
		return mvc.View(fs, "root/view.html"), nil
	})
}
