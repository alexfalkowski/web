package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewRoot returns the full-page and partial views bound to the root templates.
func NewRoot(server *mvc.Server) (*mvc.View, *mvc.View) {
	return server.NewViewPair("root/view/root.tmpl")
}
