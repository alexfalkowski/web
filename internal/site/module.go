package site

import (
	"github.com/alexfalkowski/web/internal/site/books"
	"github.com/alexfalkowski/web/internal/site/meta"
	"github.com/alexfalkowski/web/internal/site/robots"
	"github.com/alexfalkowski/web/internal/site/root"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	meta.Module,
	books.Module,
	robots.Module,
	root.Module,
	fx.Provide(NewFileSystem),
	fx.Provide(NewLayout),
)
