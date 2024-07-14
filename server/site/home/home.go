package home

import (
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register home.
func Register(fs fs.FS) {
	mvc.Route("GET /home", mvc.NewSuccessView(mvc.ParseTemplate(fs, "home/success.html")), mvc.NoController)
}
