package books

import (
	"cmp"
	"context"
	"embed"
	"net/http"
	"slices"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/go-service/runtime"
	"gopkg.in/yaml.v3"
)

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

// Register books.
func Register(fs embed.FS) {
	db = fs

	mvc.Route("GET /books", mvc.NewSuccessView(mvc.ParseTemplate(fs, "books/success.html")), controller)
}

func controller(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*Model, error) {
	d, err := db.ReadFile("books/db.yaml")
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
