# AGENTS.md

This repository is a small Go HTTP service that serves a website (templates + content embedded into a single binary) and includes a Ruby acceptance test harness.

## Quick start

### 1) Ensure the `bin/` submodule is present

This repo uses a git submodule at `bin/` (see `.gitmodules`). Many `make` targets rely on scripts/make includes from that submodule.

```sh
make submodule
```

### 2) Install deps

The top-level Make targets install both Go and Ruby deps (Ruby deps live under `test/`).

```sh
make dep
```

### 3) Run tests

Go unit/spec tests:

```sh
go test ./...
# or
make specs
```

Ruby acceptance tests (Cucumber):

```sh
make features
```

## Essential commands

All commands are driven through `make` (see `Makefile` and `bin/build/make/*.mak`). List everything available:

```sh
make help
```

Common workflows observed in this repo:

### Dependency management

- `make dep` — download/tidy/vendor Go modules + install Ruby gems under `test/`
- `make clean` — clean downloaded deps via `bin/build/go/clean`
- `make clean-dep` — clear Go caches + `bundler clean`

### Build / run

- `make build` — build release binary (outputs `./web`)
- `make dev` — run in dev mode using `air` and config `file:test/.config/server.yml`

If you built with `make build`, run the server command:

```sh
./web server
```

### Lint / format

- `make lint` — Go lint (field alignment + golangci-lint with `--build-tags features`) + Ruby lint (Rubocop in `test/`)
- `make fix-lint` — auto-fix lint where possible
- `make format` — `go fmt ./...` + Rubocop auto-format in `test/`

### Tests

- `make specs` — Go tests via `gotestsum` with race + coverage (uses vendored modules: `-mod vendor`)
- `make features` — builds a test binary (`make build-test`) then runs Cucumber in `test/`
- `make benchmarks` — runs Cucumber benchmarks (`@benchmark`-tagged)

### Security / scanning

- `make sec` — `govulncheck -test ./...`
- `make trivy-repo` — repository scan (invoked in CI)

### Coverage

- `make coverage` — generates HTML + func coverage (merges/sanitizes coverage under `test/reports/`)

## Project layout

Observed structure (matches the `golang-standards/project-layout` layout referenced in `README.md`):

- `main.go` — CLI entrypoint; registers the `server` command (`internal/cmd.RegisterServer`)
- `internal/` — application code (DI modules + MVC features)
  - `internal/cmd/` — assembles the server module graph (see `internal/cmd/module.go`)
  - `internal/config/` — service config wrapper around `go-service` base config (`internal/config/config.go`)
  - `internal/health/` — health check registrations and observers
  - `internal/site/` — website features (templates/assets embedded via `go:embed`)
- `test/` — Ruby acceptance tests (Cucumber) + a small Ruby HTTP client library
- `bin/` — git submodule providing build tooling and helper scripts

## Architecture & code patterns

### Dependency injection via `go-service/v2/di`

The server is assembled via DI modules:

- Top-level module: `internal/cmd.Module` (`internal/cmd/module.go`)
- Site features are composed in `internal/site.Module` (`internal/site/module.go`)

Pattern used throughout:

- Feature package exposes `var Module = di.Module(...)`.
- Constructors are registered via `di.Constructor(...)`.
- Route registration is commonly wired via `di.Register(Register)` where `Register(...)` calls into the MVC router.

### MVC routing + controllers

Routing is done via `go-service/v2/net/http/mvc`.

Typical route packages register both full and partial renders:

- Root (`/`): `internal/site/root/route/route.go`
- Books (`/books`): `internal/site/books/route/route.go`

Pattern:

- `mvc.Get("/path", controller.NewX(..., fullView))`
- `mvc.Put("/path", controller.NewX(..., partialView))`

Controllers typically return `mvc.Controller[T]` functions of the form:

```go
func(_ context.Context) (*mvc.View, *T, error)
```

Views are usually created in `view` packages via `mvc.NewViewPair("...tmpl")`.

### Embedded assets (templates + YAML)

`internal/site/site.go` embeds templates and YAML data with `go:embed`:

- Layout templates: `internal/site/root/layout/*.tmpl`
- Page templates: `internal/site/root/view/*.tmpl`, `internal/site/books/view/*.tmpl`
- Books data: `internal/site/books/repository/books.yaml`

`site.NewLayout()` configures full vs partial layout roots (`root/layout/full.tmpl` and `root/layout/partial.tmpl`).

### Static assets

`robots.txt` is served via MVC static file support:

- `internal/site/robots/robots.go` registers `mvc.StaticFile("/robots.txt", "robots/robots.txt")`

## Testing

### Go

- `go test ./...` runs standard Go tests.
- `make specs` uses `gotestsum` with `-race` and produces coverage artifacts under `test/reports/`.

Note: there is a build tag used for feature tests (`//go:build features` in `main_test.go`). The tooling in this repo frequently runs with `--build-tags features` (for example, golangci-lint in `bin/build/make/http.mak`).

### Ruby acceptance tests

Acceptance tests live under `test/features/` and are run with Cucumber.

The test harness uses a small Ruby client (`test/lib/web.rb`, `test/lib/web/v1/http.rb`) and a Nonnative process config (`test/nonnative.yml`) which starts the built Go binary with:

- command: `server`
- parameters: `-i file:.config/server.yml`
- base URL: `http://localhost:11000`

To work on the acceptance tests directly:

```sh
make -C test dep
make -C test features
```

## Configuration

Local/dev configuration for `make dev` is read from:

- `test/.config/server.yml`

It binds HTTP to `tcp://:11000` (see `transport.http.address` in that file) and configures health + telemetry settings.

The CLI supports passing config inputs (for example `-i file:test/.config/server.yml` as used by `make dev`).

## CI notes (CircleCI)

CircleCI runs (see `.circleci/config.yml`):

- `make clean && make dep`
- `make lint`
- `make sec` and `make trivy-repo`
- `make features` and `make benchmarks`
- `make analyse` and `make coverage`

If an agent changes Go dependencies (`go.sum`), there is a helper script `bin/build/go/clean` that triggers cache clean + `make dep clean-dep clean-lint` when `go.sum` differs from `master`.

## Gotchas / non-obvious details

- **`bin/` is a submodule**: missing or stale submodule checkout will break `make` includes and helper scripts.
- **Tooling expects external binaries**: `make dev` uses `air`; Go lint uses `golangci-lint` (wrapper script only runs it if installed); security uses `govulncheck`; Ruby tasks use `bundler`, `rubocop`, `cucumber`.
- **Field alignment scope**: `.gofa` contains `internal`, and `bin/build/go/fa` uses it to decide which packages `fieldalignment` runs against.
