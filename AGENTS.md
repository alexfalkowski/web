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
- `make codecov-upload`

`make codecov-upload` is CI upload behavior, not a read-only local validation
step.

The non-`master` workflow also runs:

- `make platform=amd64 test-docker`
- `make platform=arm64 test-docker`
- `make sync push`

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
- This repository consumes shared Make targets from the `bin/` submodule. If a
  one-command local CI preflight target is needed, it should be added to the
  shared `bin` Make fragments rather than as a service-local target here. Do not
  flag the lack of a root `verify`/`ci-checks` target as a feature gap by
  default.
- CircleCI's `version` job runs the external `package` command from the
  `alexfalkowski/release` image. That release image owns GoReleaser config
  validation before release packaging. Do not flag the absence of a separate
  repository-local GoReleaser config validation job as a project gap by default
  unless there is concrete evidence that the release image no longer validates
  `.goreleaser.yml`, or that this repository has explicitly decided to own a
  pre-release GoReleaser check locally.
- The Ruby code under `test/` is a local feature-test harness, not production
  service code. Ruby runtime selection for this harness is owned by the shared
  repository tooling and CI images. Do not flag the absence of a
  repository-local `.ruby-version`, `.tool-versions`, `mise.toml`, or Gemfile
  `ruby` directive as a project gap by default unless there is concrete evidence
  that the current workflow no longer supplies the expected runtime, or that this
  repository has explicitly decided to own Ruby version selection locally for the
  test harness.
- The Ruby code under `test/` is a local feature-test harness, not production
  service code. Fixed localhost endpoints in `test/lib/**`, `test/nonnative.yml`,
  and related feature helpers are intentional local harness assumptions unless
  there is concrete evidence of current workflow breakage. Do not flag the lack
  of environment-configurable HTTP or observability endpoints as a feature gap by
  default.

## Gotchas

- `bin/` is a required submodule; stale or missing checkout breaks many targets.
- Tooling expects external binaries such as `air`, `golangci-lint`,
  `govulncheck`, `bundler`, `rubocop`, and `cucumber`.
- Field alignment only targets packages listed in `.gofa` (`internal` here).
