package books

import (
	"cmp"
	"context"
	"embed"
	"io/fs"
	"net/http"
	"slices"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"gopkg.in/yaml.v3"
)

//go:embed db.yaml
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
func View(fs fs.FS) *mvc.View {
	return mvc.NewSuccessView(mvc.ParseTemplate(fs, "books/success.html"))
}

// Controller for books.
func Controller(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*Model, error) {
	d, err := db.ReadFile("db.yaml")
	if err != nil {
		return nil, err
	}

	var m Model
	ptr := &m

	if err := yaml.Unmarshal(d, ptr); err != nil {
		return nil, err
	}

	slices.SortFunc(ptr.Books, func(a, b *Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	return ptr, nil
}
