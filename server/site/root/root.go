package root

import (
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Path for root.
func Path() string {
	return "GET /"
}

// View for root.
func View(fs fs.FS) *mvc.View {
	return mvc.NewSuccessView(mvc.ParseTemplate(fs, "root/success.html"))
}
