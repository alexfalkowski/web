package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewBooks returns the full-page and partial views bound to the books
// templates.
func NewBooks() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("books/view/books.tmpl")
}
