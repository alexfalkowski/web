package site

import (
	"embed"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

//go:embed **/*.tmpl
//go:embed **/*.yaml
//go:embed **/*.txt
var fs embed.FS

// NewFS for site.
func NewFS() *embed.FS {
	return &fs
}

// NewPatterns for site.
func NewPatterns() mvc.Patterns {
	return mvc.Patterns{"**/*.tmpl"}
}
