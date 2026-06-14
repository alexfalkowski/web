package route

import (
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/books/controller"
	"github.com/alexfalkowski/web/internal/site/books/repository"
	"github.com/alexfalkowski/web/internal/site/books/view"
)

// Register installs the GET and PUT /books routes on the global MVC router.
func Register(repo repository.Repository) {
	view, partialView := view.NewBooks()

	mvc.Get("/books", controller.NewBooks(repo, view))
	mvc.Put("/books", controller.NewBooks(repo, partialView))
}
