package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// HomeModel for HTTP.
type HomeModel struct {
	Description string
}

// HomeRoute registers routes for home.
func HomeRoute() {
	v := mvc.NewView(mvc.ParseTemplate(Views, "home.tmpl.html"), nil)

	mvc.Route("GET /home", v, func(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*HomeModel, error) {
		m := &HomeModel{
			Description: `Lean thinking encourages teams to visualize, manage, and continuously optimize their flow. If the goal is to
						  deliver value to the customer at a sustainably fast pace, then we must have control over how work gets done; we
						  must manage flow.`,
		}

		return m, nil
	})
}
