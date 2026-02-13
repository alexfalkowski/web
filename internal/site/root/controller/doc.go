// Package controller defines MVC controllers for the root (home) feature.
//
// Controllers in this package are responsible for assembling the model data
// required by the root views. They are intentionally small and focused: they
// populate the root view model with shared site metadata and return it alongside
// the view used for rendering.
//
// # Public API
//
// NewRoot constructs an MVC controller that populates the root model with the
// provided meta information and returns it with the provided view.
//
// This package is typically wired via dependency injection by the root route
// registration rather than being used directly by consumers.
package controller
