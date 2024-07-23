package books

import (
	"cmp"
	"context"
	"embed"
	"html/template"
	"net/http"
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
func Register(fs embed.FS) {
	mvc.Route("GET /books", func(_ context.Context, _ *http.Request, _ http.ResponseWriter) *mvc.Result {
		d, err := fs.ReadFile("books/db.yaml")
		runtime.Must(err)

		var m Model
		ptr := &m

		err = yaml.Unmarshal(d, ptr)
		runtime.Must(err)

		slices.SortFunc(ptr.Books, func(a, b *Book) int {
			return cmp.Compare(a.Title, b.Title)
		})

		v := template.Must(template.ParseFS(fs, "books/view.html"))
		r := mvc.NewResult(ptr, v)

		return r
	})
}
