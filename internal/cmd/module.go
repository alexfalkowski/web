package cmd

import (
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/web/internal/config"
	"github.com/alexfalkowski/web/internal/health"
	"github.com/alexfalkowski/web/internal/site"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	module.Server,
	config.Module,
	health.Module,
	site.Module,
)
