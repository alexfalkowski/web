package books

import (
	"cmp"
	"io/fs"
	"slices"

	"gopkg.in/yaml.v3"
)

type (
	// Repository for books.
	Repository interface {
		// GetBooks from storage.
		GetBooks() (*Model, error)
	}

	// YAMLRepository has books in a file.
	YAMLRepository struct {
		filesystem fs.FS
	}
)

// NewRepository for books.
func NewRepository(filesystem fs.FS) Repository {
	return &YAMLRepository{filesystem: filesystem}
}

// GetBooks from a file.
func (r *YAMLRepository) GetBooks() (*Model, error) {
	books, err := fs.ReadFile(r.filesystem, "books/books.yaml")
	if err != nil {
		return nil, err
	}

	var m Model
	ptr := &m

	err = yaml.Unmarshal(books, ptr)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(ptr.Books, func(a, b *Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	return ptr, nil
}
