package site

import (
	"embed"
	"io/fs"

	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
)

//go:embed root/layout/*.tmpl
//go:embed root/view/*.tmpl
//go:embed books/view/*.tmpl
//go:embed books/repository/*.yaml
var filesystem embed.FS

// NewFileSystem for site.
func NewFileSystem() fs.FS {
	return filesystem
}

// NewLayout for site.
func NewLayout() *mvc.Layout {
	return mvc.NewLayout("root/layout/full.tmpl", "root/layout/partial.tmpl")
}
