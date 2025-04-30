package books

import "github.com/alexfalkowski/web/internal/site/meta"

type (
	// Books for site.
	Books struct {
		*meta.Info `yaml:"-"`
		Books      []*Book `yaml:"books,omitempty"`
	}

	// Book for site.
	Book struct {
		Title string `yaml:"title,omitempty"`
		Link  string `yaml:"link,omitempty"`
	}
)
