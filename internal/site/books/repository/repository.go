package repository

import (
	"bytes"
	"cmp"
	"io/fs"
	"slices"

	"github.com/alexfalkowski/go-service/v2/encoding/yaml"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/alexfalkowski/go-service/v2/types/ptr"
	"github.com/alexfalkowski/web/internal/site/books/model"
	"github.com/alexfalkowski/web/internal/site/meta"
)

// Repository provides access to the books data used by the site.
//
// Implementations return a fully-populated books view model (including shared
// meta information) suitable for rendering by the MVC layer.
type Repository interface {
	// GetBooks returns the books view model loaded from the underlying storage.
	GetBooks() *model.Books
}

// NewRepository constructs the default books Repository implementation.
//
// The default implementation reads books data from a YAML file in the provided
// filesystem, decodes it with the provided YAML encoder, and annotates the
// resulting model with the provided meta Info.
func NewRepository(info *meta.Info, filesystem fs.FS, enc *yaml.Encoder) Repository {
	return &FileSystemRepository{info: info, filesystem: filesystem, enc: enc}
}

// FileSystemRepository is a Repository backed by an `fs.FS` containing a YAML file.
type FileSystemRepository struct {
	// info is injected meta information that is attached to the returned model.
	info *meta.Info

	// filesystem is expected to contain `books/repository/books.yaml`.
	filesystem fs.FS

	// enc decodes the YAML payload into the books model.
	enc *yaml.Encoder
}

// GetBooks loads, decodes, and returns the books view model.
//
// The books are read from `books/repository/books.yaml`, decoded as YAML, sorted
// by Title, and enriched with the repository's meta Info.
//
// This method panics if the YAML file cannot be read or decoded (via
// `runtime.Must`).
func (r *FileSystemRepository) GetBooks() *model.Books {
	db, err := fs.ReadFile(r.filesystem, "books/repository/books.yaml")
	runtime.Must(err)

	books := ptr.Zero[model.Books]()

	err = r.enc.Decode(bytes.NewBuffer(db), books)
	runtime.Must(err)

	slices.SortFunc(books.Books, func(a, b *model.Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	books.Info = r.info

	return books
}
