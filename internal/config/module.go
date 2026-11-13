package config

import (
	"github.com/alexfalkowski/go-service/v3/config"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module wires this package into the dependency graph.
var Module = signal.Module(
	signal.Constructor(config.NewConfig[Config]),
	signal.Decorate(decorateConfig),
	signal.Constructor(healthConfig),
)
