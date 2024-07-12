package v1

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

func rootRoute() {
	mvc.Route(
		"GET /",
		mvc.NewSuccessView(mvc.ParseTemplate(fs, "root.html")),
		mvc.NoController)
}
