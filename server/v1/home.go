package v1

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

func homeRoute() {
	v := mvc.NewView(mvc.ParseTemplate(fs, "home.html"), nil)

	mvc.Route("GET /home", v, mvc.NoController)
}
