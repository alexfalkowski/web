package home

import (
	"context"
	"html/template"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register home.
func Register(fs fs.FS) {
	mvc.Route("GET /home", func(_ context.Context) (*template.Template, any) {
		return mvc.View(fs, "home/view.html"), nil
	})
}
