package books

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/web/v2/internal/site/books/repository"
	"github.com/alexfalkowski/web/v2/internal/site/books/route"
)

// Module wires this package into the dependency graph.
var Module = signal.Module(
	repository.Module,
	route.Module,
)
