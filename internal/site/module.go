package site

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/web/internal/site/books"
	"github.com/alexfalkowski/web/internal/site/errors"
	"github.com/alexfalkowski/web/internal/site/favicon"
	"github.com/alexfalkowski/web/internal/site/meta"
	"github.com/alexfalkowski/web/internal/site/robots"
	"github.com/alexfalkowski/web/internal/site/root"
	"github.com/alexfalkowski/web/internal/site/security"
	"github.com/alexfalkowski/web/internal/site/sitemap"
)

// Module wires this package into the dependency graph.
var Module = signal.Module(
	meta.Module,
	books.Module,
	errors.Module,
	favicon.Module,
	robots.Module,
	root.Module,
	security.Module,
	sitemap.Module,
	signal.Constructor(NewFileSystem),
	signal.Constructor(NewCatalog),
	signal.Constructor(NewLayout),
)
