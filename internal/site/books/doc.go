// Package books wires the "books" site feature.
//
// The books feature provides the HTTP routes, controllers, views, and data access
// needed to render the books page. It is designed to be composed into the wider
// site via dependency injection.
//
// Consumers typically do not call into this package directly. Instead, they
// include `site.Module`, which includes `books.Module` and thereby:
//   - registers the books routes, and
//   - provides any supporting constructors (repository/view/controller wiring)
//     required to serve the feature.
//
// # Public API
//
// The public API of this package is `Module`, a DI module that installs the
// books feature into the application.
package books
