package root

import (
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register root.
func Register(fs fs.FS) {
	mvc.Route("GET /", mvc.NewSuccessView(mvc.ParseTemplate(fs, "root/success.html")), mvc.NoController)
}
