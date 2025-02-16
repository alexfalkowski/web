package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/web/internal/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(env.NewVersion().String())

	cmd.RegisterServer(command)

	return command
}
