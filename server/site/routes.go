package site

import (
	"embed"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/web/server/site/books"
	"github.com/alexfalkowski/web/server/site/home"
	"github.com/alexfalkowski/web/server/site/root"
)

//go:embed **/*.html
var views embed.FS

// Register for site.
func Register() {
	mvc.Route(root.Path(), root.View(views), mvc.NoController)
	mvc.Route(home.Path(), home.View(views), mvc.NoController)
	mvc.Route(books.Path(), books.View(views), books.Controller)
}
