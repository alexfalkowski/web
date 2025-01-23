package books

import (
	"bytes"
	"cmp"
	"io/fs"
	"slices"

	"github.com/alexfalkowski/go-service/encoding/yaml"
)

// Repository for books.
type Repository interface {
	// GetBooks from storage.
	GetBooks() (*Model, error)
}

// NewRepository for books.
func NewRepository(filesystem fs.FS, enc *yaml.Encoder) Repository {
	return &FSRepository{filesystem: filesystem, enc: enc}
}

// FSRepository has books in a file.
type FSRepository struct {
	filesystem fs.FS
	enc        *yaml.Encoder
}

// GetBooks from a file.
func (r *FSRepository) GetBooks() (*Model, error) {
	books, err := fs.ReadFile(r.filesystem, "books/books.yaml")
	if err != nil {
		return nil, err
	}

	var m Model
	ptr := &m

	if err := r.enc.Decode(bytes.NewBuffer(books), ptr); err != nil {
		return nil, err
	}

	slices.SortFunc(ptr.Books, func(a, b *Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	return ptr, nil
}
