// Package view provides view constructors for the root (home) feature.
//
// The views in this package are small helpers around the MVC view system. They
// create view instances bound to the root templates, typically as a pair:
//
//   - a full-page view, and
//   - a partial view intended for incremental updates (for example HTMX-style
//     fragment rendering).
//
// # Public API
//
// NewRoot returns the view pair used by the root routes.
package view
