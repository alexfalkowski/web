package books

import (
	"bytes"
	"cmp"
	"io/fs"
	"slices"

	"github.com/alexfalkowski/go-service/encoding/yaml"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/types/ptr"
	"github.com/alexfalkowski/web/internal/site/meta"
)

// Repository for books.
type Repository interface {
	// GetBooks from storage.
	GetBooks() *Model
}

// NewRepository for books.
func NewRepository(info *meta.Info, filesystem fs.FS, enc *yaml.Encoder) Repository {
	return &FileSystemRepository{info: info, filesystem: filesystem, enc: enc}
}

// FileSystemRepository has books in a file.
type FileSystemRepository struct {
	info       *meta.Info
	filesystem fs.FS
	enc        *yaml.Encoder
}

// GetBooks from a file.
func (r *FileSystemRepository) GetBooks() *Model {
	books, err := fs.ReadFile(r.filesystem, "books/books.yaml")
	runtime.Must(err)

	model := ptr.Zero[Model]()

	err = r.enc.Decode(bytes.NewBuffer(books), model)
	runtime.Must(err)

	slices.SortFunc(model.Books, func(a, b *Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	model.Info = r.info

	return model
}
