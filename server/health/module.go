package health

import (
	"github.com/alexfalkowski/go-service/health"
	"github.com/alexfalkowski/web/server/health/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(http.NewHealthObserver),
	fx.Provide(http.NewLivenessObserver),
	fx.Provide(http.NewReadinessObserver),
	fx.Provide(NewRegistrations),
	health.Module,
)
