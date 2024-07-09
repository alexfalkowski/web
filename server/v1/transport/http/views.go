package http

import (
	"embed"
)

// Views for HTTP.
//
//go:embed home.tmpl.html
var Views embed.FS
