[![CircleCI](https://circleci.com/gh/alexfalkowski/web.svg?style=svg)](https://circleci.com/gh/alexfalkowski/web)
[![codecov](https://codecov.io/gh/alexfalkowski/web/graph/badge.svg?token=S9SPVVYQAY)](https://codecov.io/gh/alexfalkowski/web)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/web)](https://goreportcard.com/report/github.com/alexfalkowski/web)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/web.svg)](https://pkg.go.dev/github.com/alexfalkowski/web)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# 🌐 Web

A small Go service that serves the website at:

- <https://web.lean-thoughts.com/>

The service is built on top of the [`mvc`](https://github.com/alexfalkowski/go-service/tree/master/net/http/mvc) package from `go-service` and ships as a single binary with templates, content, the favicon, and static site assets embedded.

## 🧭 Background

This project is an implementation playground for the ideas outlined in:

- <https://alejandrofalkowski.substack.com/p/hyperprogress>

## ✨ What it does

At a high level the service:

- serves the home page (`/`)
- serves a books page (`/books`)
- serves `robots.txt` (`/robots.txt`)
- serves the favicon (`/favicon.ico`)
- renders a custom not-found page for missing routes
- adds browser security headers to site responses
- exposes health, liveness/readiness, and metrics endpoints

The HTML templates, error templates, books YAML data, favicon image, and robots file are embedded into the binary using `go:embed`.

## 🏗️ Architecture overview

### 🗂️ Project layout

This repo follows the structure described in:

- <https://github.com/golang-standards/project-layout>

Key directories:

- `main.go`: entrypoint for the `web` binary
- `internal/`: application code (not importable from other modules)
- `test/`: acceptance/system tests and supporting Ruby test client
- `bin/` + `Makefile`: build/dev/test automation

### 🧩 Dependency injection and modules

The service is wired with dependency injection using `go-service/v2/di`. The top-level module that assembles the server is:

- `internal/cmd.Module`

It pulls in configuration, health, and site modules.

### 🛣️ MVC routing and rendering

Routing and rendering are handled using:

- `go-service/v2/net/http/mvc`

Feature modules (e.g. books/root/robots) register their routes during DI wiring.

### 📦 Embedded assets

The site package embeds:

- templates for layout, pages, and not-found errors
- the books YAML file used to render the books page
- the favicon PNG
- `robots.txt`

See:

- `internal/site/site.go`

## 🔌 Endpoints

### 📄 Pages

- `GET /` renders the home page
- `PUT /` renders a partial/fragment version of the home page (used for incremental updates)
- `GET /books` renders the books page
- `PUT /books` renders a partial/fragment version of the books page
- `GET /robots.txt` serves the robots file as a static asset
- `GET /favicon.ico` serves the browser favicon
- missing routes render a `404` not-found page

> [!NOTE]
> The `PUT` endpoints exist to support partial rendering patterns, for example HTMX-style incremental updates. The exact response shape depends on the templates/layout configured in the MVC layer.

> [!TIP]
> Use `GET` when checking complete pages in a browser and `PUT` when checking fragment rendering.

### 🫀 Health and observability

The service registers health checks and exposes standard probe endpoints:

- `/healthz` (overall health / online)
- `/livez` (liveness)
- `/readyz` (readiness)
- `/metrics` (Prometheus metrics)

Health timings are configured via the service config under the `health` section.

### 🛡️ Response headers

Site responses include browser security headers such as `Content-Security-Policy`, `X-Content-Type-Options`, `Referrer-Policy`, `X-Frame-Options`, `Permissions-Policy`, and `Strict-Transport-Security`.

> [!NOTE]
> The acceptance suite verifies these headers for the page, robots, favicon, and not-found responses.

## 🧰 Development

### ✅ Prerequisites

Install:

- [Go](https://go.dev/)
- [Ruby](https://www.ruby-lang.org/en/)
- Bundler for the Ruby test harness

Then initialize the shared `bin/` submodule and install dependencies:

```sh
make submodule
make dep
```

> [!IMPORTANT]
> Run `make submodule` before relying on Make targets. The root `Makefile` includes shared build fragments from `bin/`, so a missing or stale submodule breaks the normal workflow.

> [!WARNING]
> Some targets require external tools in addition to Go and Ruby. For example, `make dev` uses `air`, Go checks may use `gotestsum`, `golangci-lint`, and `govulncheck`, and security checks may use Trivy.

### 🧾 Useful Make targets

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

# Run the repo-defined Go specs wrapper
make specs

# Run Cucumber acceptance tests
make features

# Run Cucumber benchmark scenarios
make benchmarks
```

> [!TIP]
> `make help` is the best way to discover the current command surface because most project workflows come from the shared `bin/` Make fragments.

### 🚀 Running locally

There are two common ways to run the service:

#### 🔁 1) Dev mode

Use the dev target (recommended while iterating):

```sh
make dev
```

This runs the service with `test/.config/server.yml`.

#### 🧱 2) Build and run the binary

Build a local binary:

```sh
make build
```

Then run it:

```sh
./web server -config file:test/.config/server.yml
```

> [!IMPORTANT]
> The current config flag is `-config`; `-c` is the short form. The CLI command is `server`, registered in `internal/cmd`, and starts the HTTP server using the DI module graph.

### 🔎 Example: verifying endpoints

Once the server is running, you can verify key endpoints.

If you started the service with `make dev` or with `-config file:test/.config/server.yml`,
the HTTP server listens on `localhost:11000`.

Pages:

```sh
curl -i http://localhost:11000/
curl -i http://localhost:11000/books
curl -i http://localhost:11000/robots.txt
curl -i http://localhost:11000/favicon.ico
```

Partial renders (PUT):

```sh
curl -i -X PUT http://localhost:11000/
curl -i -X PUT http://localhost:11000/books
```

Health:

```sh
curl -i http://localhost:11000/healthz
curl -i http://localhost:11000/livez
curl -i http://localhost:11000/readyz
curl -i http://localhost:11000/metrics
```

Not found:

```sh
curl -i http://localhost:11000/not-a-real-page
```

> [!CAUTION]
> Ports, TLS, telemetry, and other server settings come from configuration. Do not treat `test/.config/server.yml` as a production configuration.

## ⚙️ Configuration

The service config model lives in:

- `internal/config.Config`

It embeds the shared base config from `go-service` and adds a `health` section.

The canonical local example is:

- `test/.config/server.yml`

The local development config includes:

```yaml
health:
  duration: 1s
  timeout: 1s
transport:
  http:
    address: tcp://:11000
```

> [!NOTE]
> The full configuration shape also includes shared `go-service` sections such as environment, telemetry, transport, and version metadata.

## 🧪 Testing

The primary behavioral checks are the Ruby/Cucumber suites:

```sh
make features
make benchmarks
```

CI also runs:

```sh
make lint
make sec
make analyse
make coverage
```

Go support checks are still available:

```sh
make specs
go test ./...
```

> [!NOTE]
> Treat `make features` and `make benchmarks` as the authoritative product behavior checks. The Go tests support build/tooling confidence but are not the main product signal.

### 💎 Ruby acceptance test client

The Ruby test helper client is in:

- `test/lib/web.rb`
- `test/lib/web/v1/http.rb`

It provides a small wrapper around HTTP calls used by the acceptance tests.

## 🎨 Style

Go code generally follows:

- <https://github.com/uber-go/guide/blob/master/style.md>

## 🚢 Changes and releases

Releases are handled through CI and GoReleaser configuration:

- `.circleci/config.yml`
- `.goreleaser.yml`

Generated changelog text is part of the GoReleaser release flow.

> [!CAUTION]
> Release, Docker publishing, deployment, and GitHub PR targets can push to external systems. Use the read-only validation targets unless you intend to publish or update remote state.

## 📜 License

See:

- `LICENSE`
