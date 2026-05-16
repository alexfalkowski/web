package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewNotFound view.
func NewNotFound() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("errors/view/not_found.tmpl")
}
