// Package model defines the data structures used by the root (home) feature.
//
// The types in this package represent the view model rendered by the root page
// templates. They are populated by the root controllers and are intentionally
// small: the root page primarily needs shared site metadata (see the meta
// package) to render things like the running version and current year.
//
// # Public API
//
// This package's public API is the set of exported types (for example Root)
// that are used by the root feature and passed to MVC views.
package model
