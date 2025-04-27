[![CircleCI](https://circleci.com/gh/alexfalkowski/web.svg?style=svg)](https://circleci.com/gh/alexfalkowski/web)
[![codecov](https://codecov.io/gh/alexfalkowski/web/graph/badge.svg?token=S9SPVVYQAY)](https://codecov.io/gh/alexfalkowski/web)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/web)](https://goreportcard.com/report/github.com/alexfalkowski/web)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/web.svg)](https://pkg.go.dev/github.com/alexfalkowski/web)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Web

A website for lean-thoughts.com

## Background

Try the ideas outlined in https://alejandrofalkowski.substack.com/p/hyperprogress

## Design

The design is heavily reliant on [mvc](https://github.com/alexfalkowski/go-service/tree/master/net/http/mvc).

### Server

The server code contains health and mvc. To get a better idea take a look at the site [layout](internal/site).

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](https://www.ruby-lang.org/en/)
- [Golang](https://go.dev/)

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Check out [CI](.circleci/config.yml).

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
