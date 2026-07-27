package repository

import (
	"cmp"
	"io/fs"
	"slices"

	"github.com/alexfalkowski/go-service/v2/encoding"
	"github.com/alexfalkowski/go-service/v2/ptr"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/alexfalkowski/web/internal/site/books/model"
	"github.com/alexfalkowski/web/internal/site/meta"
)

// Repository provides access to the books data used by the site.
//
// Implementations return a fully-populated books view model (including shared
// meta information) suitable for rendering by the MVC layer.
type Repository interface {
	// GetBooks returns the books view model loaded from the underlying storage.
	//
	// The default filesystem-backed implementation panics if the embedded books
	// YAML file cannot be read or decoded.
	GetBooks() *model.Books
}

// NewRepository constructs the default books Repository implementation.
//
// The default implementation reads books data from a YAML file in the provided
// filesystem, decodes it with the provided YAML encoder, and annotates the
// resulting model with the provided meta Info.
func NewRepository(info *meta.Info, fs fs.FS, m *encoding.Map) Repository {
	return &FileSystemRepository{info: info, fs: fs, m: m}
}

// FileSystemRepository is a Repository backed by an `fs.FS` containing a YAML file.
type FileSystemRepository struct {
	info *meta.Info
	fs   fs.FS
	m    *encoding.Map
}

// GetBooks loads, decodes, and returns the books view model.
//
// The books are read from `books/repository/books.yaml`, decoded as YAML, sorted
// by Title, and enriched with the repository's meta Info.
//
// This method panics if the YAML file cannot be read or decoded (via
// `runtime.Must`).
func (r *FileSystemRepository) GetBooks() *model.Books {
	file, err := r.fs.Open("books/repository/books.yaml")
	runtime.Must(err)
	defer file.Close()

	books := ptr.Zero[model.Books]()

	err = r.m.Get("yaml").Decode(file, books)
	runtime.Must(err)

	slices.SortFunc(books.Books, func(a, b *model.Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	books.Info = r.info
	books.Page = meta.Page{
		Title:       "Books | Lean Thoughts",
		Description: "Recommended books about lean thinking, software delivery, and team flow.",
	}

	return books
}
