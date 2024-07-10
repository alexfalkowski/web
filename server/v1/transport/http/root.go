package http

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// RootRoute registers routes for root (index).
func RootRoute() {
	v := mvc.NewView(mvc.ParseTemplate(Views, "root.tmpl.html"), nil)

	mvc.Route("GET /", v, mvc.NoController)
}
