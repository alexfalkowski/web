package cmd

import (
	"github.com/alexfalkowski/go-service/v3/module"
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/web/internal/config"
	"github.com/alexfalkowski/web/internal/health"
	"github.com/alexfalkowski/web/internal/site"
)

// Module wires this package into the dependency graph.
var Module = signal.Module(
	module.Server,
	config.Module,
	health.Module,
	site.Module,
)
