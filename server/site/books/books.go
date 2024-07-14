package books

import (
	"cmp"
	"context"
	"embed"
	"io/fs"
	"net/http"
	"slices"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/go-service/runtime"
	"gopkg.in/yaml.v3"
)

//go:embed books.yaml
var db embed.FS

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

// Path for books.
func Path() string {
	return "GET /books"
}

// View for books.
func View(views fs.FS) *mvc.View {
	return mvc.NewSuccessView(mvc.ParseTemplate(views, "books/success.html"))
}

// Controller for books.
func Controller(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*Model, error) {
	d, err := db.ReadFile("books.yaml")
	runtime.Must(err)

	var m Model
	ptr := &m

	err = yaml.Unmarshal(d, ptr)
	runtime.Must(err)

	slices.SortFunc(ptr.Books, func(a, b *Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	return ptr, nil
}
