package site

import (
	"github.com/alexfalkowski/web/site/books"
	"github.com/alexfalkowski/web/site/home"
	"github.com/alexfalkowski/web/site/robots"
	"github.com/alexfalkowski/web/site/root"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	books.Module,
	home.Module,
	robots.Module,
	root.Module,
	fx.Provide(NewFS),
	fx.Provide(NewPatterns),
)
