// Package repository provides access to the books data used by the site.
//
// The books repository is responsible for loading the books view model from
// storage. In this service the default implementation reads a YAML file from an
// embedded filesystem and decodes it into the books model.
//
// This package is designed to be wired via dependency injection (see Module in
// this package). Most callers should depend only on the Repository interface.
//
// # Public API
//
//   - Repository defines the interface used by controllers to fetch books.
//   - NewRepository constructs the default repository implementation.
//
// The concrete implementation type (for example FileSystemRepository) is
// exported only for completeness and should generally not be referenced
// directly by consumers.
package repository
