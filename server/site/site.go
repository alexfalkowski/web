package site

import (
	"embed"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/web/server/site/books"
	"github.com/alexfalkowski/web/server/site/home"
	"github.com/alexfalkowski/web/server/site/root"
)

//go:embed **/*.tmpl
//go:embed **/*.yaml
var fs embed.FS

// NewFS for site.
func NewFS() *embed.FS {
	return &fs
}

// NewPatterns for site.
func NewPatterns() mvc.Patterns {
	return mvc.Patterns{"**/*.tmpl"}
}

// Register for site.
func Register(router *mvc.Router, fs *embed.FS) {
	root.Register(router)
	home.Register(router)
	books.Register(router, fs)
}
