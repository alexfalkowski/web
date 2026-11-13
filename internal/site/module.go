package site

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/web/v2/internal/site/books"
	"github.com/alexfalkowski/web/v2/internal/site/errors"
	"github.com/alexfalkowski/web/v2/internal/site/favicon"
	"github.com/alexfalkowski/web/v2/internal/site/meta"
	"github.com/alexfalkowski/web/v2/internal/site/robots"
	"github.com/alexfalkowski/web/v2/internal/site/root"
	"github.com/alexfalkowski/web/v2/internal/site/security"
	"github.com/alexfalkowski/web/v2/internal/site/sitemap"
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
