package site

import (
	"github.com/alexfalkowski/web/internal/site/books"
	"github.com/alexfalkowski/web/internal/site/home"
	"github.com/alexfalkowski/web/internal/site/robots"
	"github.com/alexfalkowski/web/internal/site/root"
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
