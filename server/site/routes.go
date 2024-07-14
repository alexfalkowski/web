package site

import (
	"github.com/alexfalkowski/web/server/site/books"
	"github.com/alexfalkowski/web/server/site/home"
	"github.com/alexfalkowski/web/server/site/root"
)

// Register for site.
func Register() {
	root.Register(fs)
	home.Register(fs)
	books.Register(fs)
}
