package errors

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/web/internal/site/errors/route"
)

// Module wires this package into the dependency graph.
var Module = signal.Module(
	route.Module,
)
