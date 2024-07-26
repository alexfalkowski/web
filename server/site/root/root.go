package root

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register root.
func Register(fs fs.FS) {
	mvc.Route("GET /", func(_ context.Context) (*mvc.View, mvc.Model) {
		return mvc.NewView(fs, "root/view.html"), nil
	})
}
