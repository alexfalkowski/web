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
func View(fs fs.FS) *mvc.View {
	return mvc.NewSuccessView(mvc.ParseTemplate(fs, "home/success.html"))
}
