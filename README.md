[![CircleCI](https://circleci.com/gh/alexfalkowski/web.svg?style=svg)](https://circleci.com/gh/alexfalkowski/web)
[![codecov](https://codecov.io/gh/alexfalkowski/web/graph/badge.svg?token=S9SPVVYQAY)](https://codecov.io/gh/alexfalkowski/web)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/web)](https://goreportcard.com/report/github.com/alexfalkowski/web)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/web.svg)](https://pkg.go.dev/github.com/alexfalkowski/web)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Web

A small Go service that serves the website at:

- https://web.lean-thoughts.com/

The service is built on top of the [`mvc`](https://github.com/alexfalkowski/go-service/tree/master/net/http/mvc) package from `go-service` and ships as a single binary with templates and content embedded.

## Background

This project is an implementation playground for the ideas outlined in:

- https://alejandrofalkowski.substack.com/p/hyperprogress

## What it does

At a high level the service:

- serves the home page (`/`)
- serves a books page (`/books`)
- serves `robots.txt` (`/robots.txt`)
- exposes health and liveness/readiness endpoints (see below)

The HTML templates and the books YAML data are embedded into the binary using `go:embed`.

## Architecture overview

### Project layout

This repo follows the structure described in:

- https://github.com/golang-standards/project-layout

Key directories:

- `main.go`: entrypoint for the `web` binary
- `internal/`: application code (not importable from other modules)
- `test/`: acceptance/system tests and supporting Ruby test client
- `bin/` + `Makefile`: build/dev/test automation

### Dependency injection and modules

The service is wired with dependency injection using `go-service/v2/di`. The top-level module that assembles the server is:

- `internal/cmd.Module`

It pulls in configuration, health, and site modules.

### MVC routing and rendering

Routing and rendering are handled using:

- `go-service/v2/net/http/mvc`

Feature modules (e.g. books/root/robots) register their routes during DI wiring.

### Embedded assets

The site package embeds:

- templates for layout and pages
- the books YAML file used to render the books page

See:

- `internal/site/site.go`

## Endpoints

### Pages

- `GET /` renders the home page
- `PUT /` renders a partial/fragment version of the home page (used for incremental updates)
- `GET /books` renders the books page
- `PUT /books` renders a partial/fragment version of the books page
- `GET /robots.txt` serves the robots file as a static asset

> Note: the `PUT` endpoints exist to support partial rendering patterns (for example HTMX-style incremental updates). The exact response shape depends on the templates/layout configured in the MVC layer.

### Health

The service registers health checks and exposes standard probe endpoints:

- `/healthz` (overall health / online)
- `/livez` (liveness)
- `/readyz` (readiness)

Health timings are configured via the service config under the `health` section.

## Development

### Prerequisites

Install:

- [Go](https://go.dev/)
- [Ruby](https://www.ruby-lang.org/en/)

### Useful Make targets

This repo relies on `make` for a consistent developer experience.

List all available commands:

```sh
make help
```

Common workflows:

```sh
# Install dependencies (Go + Ruby)
make dep

# Run linters
make lint

# Auto-fix lint where possible
make fix-lint

# Format code
make format

# Run Go tests
go test ./...

# Run Ruby specs / acceptance tests (if you are working on the test harness)
make specs
make features
```

### Running locally

There are two common ways to run the service:

#### 1) Dev mode

Use the dev target (recommended while iterating):

```sh
make dev
```

This typically runs the service with a development configuration (see the `Makefile` includes and CI configuration for how it is invoked).

#### 2) Build and run the binary

Build a local binary:

```sh
make build
```

Then run it:

```sh
./tmp/web server
```

> The CLI command is `server`. It is registered in `internal/cmd` and starts the HTTP server using the DI module graph.

### Example: verifying endpoints

Once the server is running, you can verify key endpoints.

Pages:

```sh
curl -i http://localhost:8080/
curl -i http://localhost:8080/books
curl -i http://localhost:8080/robots.txt
```

Partial renders (PUT):

```sh
curl -i -X PUT http://localhost:8080/
curl -i -X PUT http://localhost:8080/books
```

Health:

```sh
curl -i http://localhost:8080/healthz
curl -i http://localhost:8080/livez
curl -i http://localhost:8080/readyz
```

> Ports, TLS, and other server settings are controlled by configuration. If your local environment differs, inspect the config used by `make dev` / CI.

## Configuration

The service config model lives in:

- `internal/config.Config`

It embeds the shared base config from `go-service` and adds a `health` section.

The health section looks conceptually like:

```yaml
health:
  duration: 5s
  timeout: 2s
```

> The exact full configuration file shape depends on the embedded `go-service` base config and the environment in which you run the binary. For canonical examples, check the CI configuration and any config files used by `make dev`.

## Testing

### Go

Run all Go tests:

```sh
go test ./...
```

### Ruby acceptance test client

The Ruby test helper client is in:

- `test/lib/web.rb`
- `test/lib/web/v1/http.rb`

It provides a small wrapper around HTTP calls used by the acceptance tests.

## Style

Go code generally follows:

- https://github.com/uber-go/guide/blob/master/style.md

## Changes / releases

See:

- `CHANGELOG.md`

## License

See:

- `LICENSE`
