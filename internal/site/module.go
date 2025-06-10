package site

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/web/internal/site/books"
	"github.com/alexfalkowski/web/internal/site/meta"
	"github.com/alexfalkowski/web/internal/site/robots"
	"github.com/alexfalkowski/web/internal/site/root"
)

// Module for fx.
var Module = di.Module(
	meta.Module,
	books.Module,
	robots.Module,
	root.Module,
	di.Constructor(NewFileSystem),
	di.Constructor(NewLayout),
)
