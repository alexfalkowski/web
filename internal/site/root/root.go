package root

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/meta"
)

// Model for root.
type Model struct {
	*meta.Info
}

// Register root.
func Register(info *meta.Info) {
	mvc.Route("GET /", func(_ context.Context) (*mvc.View, *Model, error) {
		m := &Model{Info: info}

		return mvc.NewView("root/root.tmpl"), m, nil
	})

	mvc.Route("PUT /", func(_ context.Context) (*mvc.View, *Model, error) {
		m := &Model{Info: info}

		return mvc.NewPartialView("root/root.tmpl"), m, nil
	})
}
