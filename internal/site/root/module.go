package root

import (
	"github.com/alexfalkowski/web/internal/site/root/route"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	route.Module,
)
