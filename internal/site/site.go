package site

import (
	"embed"

	"github.com/alexfalkowski/go-service/v2/i18n"
	"github.com/alexfalkowski/go-service/v2/io"
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"golang.org/x/text/language"
)

//go:embed root/layout/*.tmpl
//go:embed root/view/*.tmpl
//go:embed errors/view/*.tmpl
//go:embed books/view/*.tmpl
//go:embed books/repository/*.yaml
//go:embed favicon/*.png
//go:embed robots/*.txt
//go:embed sitemap/*.xml
//go:embed locales/*.yaml
var filesystem embed.FS

// NewFileSystem returns an `io.FS` backed by the site's embedded assets.
//
// The returned filesystem contains the templates and data files embedded via the
// `go:embed` directives in this package. It is intended to be injected into
// downstream components (for example repositories) that need access to those
// assets at runtime.
func NewFileSystem() io.FS {
	return filesystem
}

// NewCatalog loads the embedded site messages with English as the fallback.
// Catalog read, decode, and template syntax errors prevent service startup.
func NewCatalog(fs io.FS) (*i18n.Catalog, error) {
	return i18n.NewCatalog(language.English, fs, "locales/*.yaml")
}

// NewLayout constructs the site's MVC layout.
//
// It returns a layout configured with the full-page and partial template roots
// used by the site, enabling both full renders and fragment/partial renders.
func NewLayout() *mvc.Layout {
	return mvc.NewLayout("root/layout/full.tmpl", "root/layout/partial.tmpl")
}
