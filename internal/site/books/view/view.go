package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewBooks view.
func NewBooks() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("books/view/books.tmpl")
}
