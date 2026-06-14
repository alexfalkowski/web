package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewRoot returns the full-page and partial views bound to the root templates.
func NewRoot() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("root/view/root.tmpl")
}
