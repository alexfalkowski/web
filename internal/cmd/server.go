package cmd

import (
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// RegisterServer registers the service's "server" command on the provided CLI commander.
//
// The registered command starts the web server using the dependency-injection graph
// defined by `Module` in this package.
//
// This function is intended to be called by the `main` package during application
// initialization.
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start web server", Module)

	cmd.AddInput(strings.Empty)
}
