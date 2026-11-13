package health

import "github.com/alexfalkowski/go-signal/v2"

// Module wires this package into the dependency graph.
var Module = signal.Module(
	signal.Register(register),
	signal.Register(httpHealthObserver),
	signal.Register(httpLivenessObserver),
	signal.Register(httpReadinessObserver),
)
