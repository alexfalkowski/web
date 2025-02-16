package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/web/internal/config"
	"github.com/alexfalkowski/web/internal/health"
	"github.com/alexfalkowski/web/internal/site"
)

// RegisterServer  for cmd.
func RegisterServer(command *cmd.Command) {
	flags := flags.NewFlagSet("server")
	flags.AddInput("env:WEB_CONFIG_FILE")

	command.AddServer("server", "Start web server", flags,
		module.Module, debug.Module, feature.Module,
		telemetry.Module, transport.Module, health.Module,
		config.Module, cmd.Module, site.Module,
	)
}
