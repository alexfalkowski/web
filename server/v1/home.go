package v1

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

func homeRoute() {
	mvc.Route(
		"GET /home",
		mvc.NewSuccessView(mvc.ParseTemplate(fs, "home.html")),
		mvc.NoController,
	)
}
