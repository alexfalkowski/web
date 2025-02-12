package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/web/internal/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:WEB_CONFIG_FILE")
	c.AddServer("server", "Start web server", cmd.ServerOptions...)

	return c
}
