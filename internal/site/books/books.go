package books

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register books.
func Register(repo Repository) error {
	model := repo.GetBooks()

	mvc.Route("GET /books", func(_ context.Context) (*mvc.View, *Model, error) {
		return mvc.NewView("books/books.tmpl"), model, nil
	})

	return nil
}
