package root

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/web/v2/internal/site/root/route"
)

// Module wires this package into the dependency graph.
var Module = signal.Module(
	route.Module,
)
