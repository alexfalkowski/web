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

// NewFileSystem returns an `fs.FS` backed by the site's embedded assets.
//
// The returned filesystem contains the templates and data files embedded via the
// `go:embed` directives in this package. It is intended to be injected into
// downstream components (for example repositories) that need access to those
// assets at runtime.
func NewFileSystem() fs.FS {
	return filesystem
}

// NewLayout constructs the site's MVC layout.
//
// It returns a layout configured with the full-page and partial template roots
// used by the site, enabling both full renders and fragment/partial renders.
func NewLayout() *mvc.Layout {
	return mvc.NewLayout("root/layout/full.tmpl", "root/layout/partial.tmpl")
}
