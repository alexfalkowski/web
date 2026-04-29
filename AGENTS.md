# AGENTS.md

Repo-specific guidance for this repository.

## Shared skill

Use the shared `coding-standards` skill from `bin/skills/coding-standards` for code changes, bug fixes, refactors, reviews, tests, linting, documentation, PR summaries, commits, Makefile changes, CI validation, and verification.

## What this repo is

- Small Go HTTP service serving a website from embedded templates and content.
- Ruby acceptance test harness under `test/`.
- Build/test automation comes from the `bin/` submodule and root `Makefile`.

## First steps

1. Ensure the `bin/` submodule is present:

```sh
make submodule
```

2. Install dependencies:

```sh
make dep
```

3. Prefer repo-defined commands:

```sh
make help
```

## Commands that matter

- `make dep` installs Go deps and Ruby gems.
- `make lint` runs Go lint plus Ruby lint.
- `make build` builds the binary as `./web`.
- `make dev` runs the service with `file:test/.config/server.yml`.
- `make features` runs the primary Cucumber acceptance suite.
- `make benchmarks` runs the `@benchmark` Cucumber scenarios.
- `make sec` runs security checks.
- `make coverage` generates coverage artifacts under `test/reports/`.

If you build locally, start the service with:

```sh
./web server
```

## Testing policy

- The intentional product test strategy is the Ruby/Cucumber suite in
  `test/features/`.
- Treat `make features` and `make benchmarks` as the authoritative behavioral
  checks.
- `go test ./...` and `make specs` exist for tooling/build support, but they are
  not the primary product test signal in this repo.

## Architecture

- CLI entrypoint: `main.go`
- Server wiring: `internal/cmd`
- Config wrapper: `internal/config`
- Health registrations and observers: `internal/health`
- Site feature composition: `internal/site`

The service is assembled with `go-service/v2/di`.

Common pattern:

- Each feature exposes `var Module = di.Module(...)`.
- Constructors use `di.Constructor(...)`.
- Route registration usually uses `di.Register(Register)`.

Routing uses `go-service/v2/net/http/mvc`.

Common route pattern:

- `mvc.Get("/path", controller.NewX(..., fullView))`
- `mvc.Put("/path", controller.NewX(..., partialView))`

Embedded assets live in `internal/site/site.go` and include:

- layout templates
- page templates
- `internal/site/books/repository/books.yaml`

`robots.txt` is served from `internal/site/robots`.

## Config and local runtime

- Dev config: `test/.config/server.yml`
- Dev/test HTTP address: `tcp://:11000`
- `make dev` runs with `-i file:test/.config/server.yml`

The acceptance harness uses `test/nonnative.yml` and starts:

- command: `server`
- parameters: `-i file:.config/server.yml`
- base URL: `http://localhost:11000`

## CI truth

CircleCI is the source of truth for required checks. The main pipeline runs:

- `make clean && make dep`
- `make lint`
- `make sec`
- `make features`
- `make benchmarks`
- `make analyse`
- `make coverage`

## Gotchas

- `bin/` is a required submodule; stale or missing checkout breaks many targets.
- Tooling expects external binaries such as `air`, `golangci-lint`,
  `govulncheck`, `bundler`, `rubocop`, and `cucumber`.
- Field alignment only targets packages listed in `.gofa` (`internal` here).
