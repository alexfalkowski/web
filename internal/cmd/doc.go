// Package cmd wires up the service's CLI commands.
//
// This package is responsible for registering the commands exposed by the
// `web` binary and connecting them to the dependency-injection graph used to
// construct and run the server.
//
// # Public API
//
// The public surface of this package is intended to be small and stable:
//   - RegisterServer registers the "server" command on a provided CLI commander.
//   - Module is the DI module that assembles all dependencies required to run
//     the server command.
package cmd
