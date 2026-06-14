# AGENTS.md

Repo-specific guidance for this repository.

## Shared skills

This repository uses the shared skills from `bin/skills/`. Read
`bin/AGENTS.md` for the canonical shared skill list and use the smallest
matching skill for the task.

## What this repo is

- Small Go HTTP service serving a website from embedded templates and content.
- Ruby acceptance test harness under `test/`.
- Build/test automation comes from the `bin/` submodule and root `Makefile`.

## First steps

1. Ensure the `bin/` submodule is present before using Make targets:

```sh
git submodule sync
git submodule update --init
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
./web server -config file:test/.config/server.yml
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
- `make dev` runs with `-config file:test/.config/server.yml`

The acceptance harness uses `test/nonnative.yml` and starts:

- command: `server`
- parameters: `-config file:.config/server.yml`
- base URL: `http://localhost:11000`

## CI truth

CircleCI is the source of truth for required checks. The main service build job runs:

- `make clean && make dep`
- `make lint`
- `make sec`
- `make features`
- `make benchmarks`
- `make analyse`
- `make coverage`

The non-`master` workflow also runs:

- `make platform=amd64 test-docker`
- `make platform=arm64 test-docker`

The `master` workflow also runs versioning, Docker release/manifest, and deploy
jobs.

## Intentional design choices

- The `/healthz` observer intentionally uses `go-health/v2`'s default
  `server.NewOnlineRegistration` connectivity check. That check reaches public
  connectivity URLs by default; do not flag the lack of configurable online
  health URLs as an issue unless the task is explicitly about changing health
  check semantics.
- Site metadata such as the footer year is intentionally computed at startup
  and shared through DI. Do not flag year rollover staleness unless the task is
  explicitly about changing metadata freshness.

## Gotchas

- `bin/` is a required submodule; stale or missing checkout breaks many targets.
- Tooling expects external binaries such as `air`, `golangci-lint`,
  `govulncheck`, `bundler`, `rubocop`, and `cucumber`.
- Field alignment only targets packages listed in `.gofa` (`internal` here).
