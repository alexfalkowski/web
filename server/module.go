package server

import (
	"github.com/alexfalkowski/web/server/site"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	site.Module,
)
