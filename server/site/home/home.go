package home

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register home.
func Register(fs fs.FS) {
	mvc.Route("GET /home", func(_ context.Context) (*mvc.View, mvc.Model) {
		return mvc.NewView(fs, "home/view.html"), nil
	})
}
