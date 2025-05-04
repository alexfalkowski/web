package controller

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/books/model"
	"github.com/alexfalkowski/web/internal/site/books/repository"
)

// NewRoot controller.
func NewBooks(repo repository.Repository, view *mvc.View) mvc.Controller[model.Books] {
	return func(_ context.Context) (*mvc.View, *model.Books, error) {
		model := repo.GetBooks()

		return view, model, nil
	}
}
