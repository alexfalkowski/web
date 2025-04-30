package books

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register books.
func Register(repo Repository) error {
	booksFullView, booksPartialView := mvc.NewViewPair("books/books.tmpl")

	mvc.Get("/books", func(_ context.Context) (*mvc.View, *Books, error) {
		model := repo.GetBooks()

		return booksFullView, model, nil
	})

	mvc.Put("/books", func(_ context.Context) (*mvc.View, *Books, error) {
		model := repo.GetBooks()

		return booksPartialView, model, nil
	})

	return nil
}
