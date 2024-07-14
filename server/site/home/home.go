package home

import (
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Path for home.
func Path() string {
	return "GET /home"
}

// View for home.
func View(views fs.FS) *mvc.View {
	return mvc.NewSuccessView(mvc.ParseTemplate(views, "home/success.html"))
}
