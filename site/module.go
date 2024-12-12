package site

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewFS),
	fx.Provide(NewPatterns),
	fx.Invoke(Register),
)
