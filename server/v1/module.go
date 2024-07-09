package v1

import (
	v1 "github.com/alexfalkowski/web/server/v1/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Invoke(v1.HomeRoute),
)
