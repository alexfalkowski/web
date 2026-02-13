// Package site wires the HTTP site (routes, views, and embedded assets).
//
// The site package provides constructors for shared site dependencies and bundles
// the sub-features that make up the public-facing website (for example root,
// books, and robots). It also embeds the templates and content files required by
// those features so the service can be shipped as a single binary.
//
// This package is intended to be used via dependency injection. In practice,
// consumers include the `cmd` package's DI module, which pulls in `site.Module`
// and thereby registers all routes and supporting dependencies.
//
// # Public API
//
//   - Module assembles and registers all site submodules and shared constructors.
//   - NewFileSystem returns an fs.FS backed by embedded site assets.
//   - NewLayout returns the MVC layout used by the site templates.
package site
