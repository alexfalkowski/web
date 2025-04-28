package books

import "github.com/alexfalkowski/web/internal/site/meta"

type (
	// Model for books.
	Model struct {
		*meta.Info `yaml:"-"`
		Books      []*Book `yaml:"books,omitempty"`
	}

	// Book for books.
	Book struct {
		Title string `yaml:"title,omitempty"`
		Link  string `yaml:"link,omitempty"`
	}
)
