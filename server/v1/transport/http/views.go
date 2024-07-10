package http

import (
	"embed"
)

// Views for HTTP.
//
//go:embed root.tmpl.html
//go:embed home.tmpl.html
//go:embed books.tmpl.html
var Views embed.FS
