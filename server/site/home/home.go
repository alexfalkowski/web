package home

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register home.
func Register(fs fs.FS) {
	mvc.Route("GET /home", func(_ context.Context) *mvc.Result {
		v := mvc.View(fs, "home/view.html")
		r := mvc.NewResult(nil, v)

		return r
	})
}
