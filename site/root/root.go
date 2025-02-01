package root

import (
	"context"
	"time"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Model for root.
type Model struct {
	Version env.Version
	Year    int
}

// Register root.
func Register(version env.Version) {
	mvc.Route("GET /", func(_ context.Context) (mvc.View, *Model, error) {
		m := &Model{
			Year:    time.Now().Year(),
			Version: version,
		}

		return mvc.View("root.tmpl"), m, nil
	})
}
