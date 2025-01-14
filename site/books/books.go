package books

import (
	"cmp"
	"context"
	"embed"
	"slices"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/go-service/runtime"
	"gopkg.in/yaml.v3"
)

type (
	// Model for books.
	Model struct {
		Books []*Book `yaml:"books,omitempty"`
	}

	// Book for books.
	Book struct {
		Title string `yaml:"title,omitempty"`
		Link  string `yaml:"link,omitempty"`
	}
)

// Register books.
func Register(router *mvc.Router, fs *embed.FS) {
	router.Route("GET /books", func(_ context.Context) (mvc.View, mvc.Model) {
		books, err := fs.ReadFile("books/books.yaml")
		runtime.Must(err)

		var m Model
		ptr := &m

		err = yaml.Unmarshal(books, ptr)
		runtime.Must(err)

		slices.SortFunc(ptr.Books, func(a, b *Book) int {
			return cmp.Compare(a.Title, b.Title)
		})

		return mvc.View("books.tmpl"), ptr
	})
}
