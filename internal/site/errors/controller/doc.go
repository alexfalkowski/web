// Package controller defines MVC controllers for the error feature.
//
// Controllers in this package are responsible for assembling the model data
// required by error views. They choose between full-page and partial views based
// on the request method.
//
// # Public API
//
// NewNotFound constructs an MVC not-found controller that populates the error
// model with shared site metadata and returns the appropriate view.
//
// This package is typically wired via dependency injection by the error route
// registration rather than being used directly by consumers.
package controller
