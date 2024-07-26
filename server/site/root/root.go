package root

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register root.
func Register(fs fs.FS) {
	mvc.Route("GET /", func(_ context.Context) *mvc.Result {
		v := mvc.View(fs, "root/view.html")
		r := mvc.NewResult(nil, v)

		return r
	})
}
