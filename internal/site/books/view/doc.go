// Package view provides view constructors for the books feature.
//
// The views in this package are thin helpers around the MVC view system. They
// create view instances bound to the books templates, typically as a pair:
//
//   - a full-page view, and
//   - a partial view intended for incremental updates (for example HTMX-style
//     fragment rendering).
//
// # Public API
//
// NewBooks returns the view pair used by the books routes.
package view
