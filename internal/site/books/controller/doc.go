// Package controller defines MVC controllers for the books feature.
//
// The controllers in this package are responsible for assembling the model
// data required by the books views. They are intentionally thin: they pull
// data from the books repository and return it alongside the appropriate view
// for rendering.
//
// # Public API
//
// NewBooks constructs an MVC controller that loads the books model from the
// provided repository and returns it with the provided view.
//
// This package is typically wired via dependency injection by the books route
// registration rather than being used directly by consumers.
package controller
