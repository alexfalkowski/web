package books

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register books.
func Register(repo Repository) error {
	mvc.Get("/books", func(_ context.Context) (*mvc.View, *Model, error) {
		model := repo.GetBooks()

		return mvc.NewView("books/books.tmpl"), model, nil
	})

	mvc.Put("/books", func(_ context.Context) (*mvc.View, *Model, error) {
		model := repo.GetBooks()

		return mvc.NewPartialView("books/books.tmpl"), model, nil
	})

	return nil
}
