package site

import (
	"embed"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/web/site/books"
	"github.com/alexfalkowski/web/site/home"
	"github.com/alexfalkowski/web/site/robots"
	"github.com/alexfalkowski/web/site/root"
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

// Register for site.
func Register(router *mvc.Router, fs *embed.FS, version env.Version) {
	root.Register(router, version)
	home.Register(router)
	books.Register(router, fs)
	robots.Register(router)
}
