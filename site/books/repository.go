package books

import (
	"cmp"
	"embed"
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
		fs *embed.FS
	}
)

// NewRepository for books.
func NewRepository(fs *embed.FS) Repository {
	return &YAMLRepository{fs: fs}
}

// GetBooks from a file.
func (r *YAMLRepository) GetBooks() (*Model, error) {
	books, err := r.fs.ReadFile("books/books.yaml")
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
