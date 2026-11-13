package meta

import "github.com/alexfalkowski/go-signal/v2"

// Module wires this package into the dependency graph.
var Module = signal.Module(
	signal.Constructor(NewYear),
	signal.Constructor(NewInfo),
)
