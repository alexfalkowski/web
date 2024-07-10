package v1

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

func rootRoute() {
	v := mvc.NewView(mvc.ParseTemplate(fs, "root.html"), nil)

	mvc.Route("GET /", v, mvc.NoController)
}
