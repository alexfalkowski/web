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
func View(views fs.FS) *mvc.View {
	return mvc.NewSuccessView(mvc.ParseTemplate(views, "root/success.html"))
}
