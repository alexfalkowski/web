package site

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/web/internal/site/books"
	"github.com/alexfalkowski/web/internal/site/errors"
	"github.com/alexfalkowski/web/internal/site/favicon"
	"github.com/alexfalkowski/web/internal/site/meta"
	"github.com/alexfalkowski/web/internal/site/robots"
	"github.com/alexfalkowski/web/internal/site/root"
	"github.com/alexfalkowski/web/internal/site/security"
)

// Module wires this package into the dependency graph.
var Module = di.Module(
	meta.Module,
	books.Module,
	errors.Module,
	favicon.Module,
	robots.Module,
	root.Module,
	security.Module,
	di.Constructor(NewFileSystem),
	di.Constructor(NewLayout),
)
