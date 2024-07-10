package http

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// HomeRoute registers routes for home.
func HomeRoute() {
	v := mvc.NewView(mvc.ParseTemplate(Views, "home.tmpl.html"), nil)

	mvc.Route("GET /home", v, mvc.NoController)
}
