package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewNotFound returns the full-page and partial views bound to the not-found
// templates.
func NewNotFound(server *mvc.Server) (*mvc.View, *mvc.View) {
	return server.NewViewPair("errors/view/not_found.tmpl")
}
