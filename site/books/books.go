package books

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register books.
func Register(router *mvc.Router, repo Repository) error {
	model, err := repo.GetBooks()
	if err != nil {
		return err
	}

	router.Route("GET /books", func(_ context.Context) (mvc.View, mvc.Model) {
		return mvc.View("books.tmpl"), model
	})

	return nil
}
