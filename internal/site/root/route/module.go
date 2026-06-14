package route

import "github.com/alexfalkowski/go-service/v2/di"

// Module wires this package into the dependency graph.
var Module = di.Module(
	di.Register(Register),
)
