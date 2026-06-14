package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewNotFound returns the full-page and partial views bound to the not-found
// templates.
func NewNotFound() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("errors/view/not_found.tmpl")
}
