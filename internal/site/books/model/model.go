package model

import "github.com/alexfalkowski/web/internal/site/meta"

type (
	// Books is the view model rendered by the books page.
	Books struct {
		Info  *meta.Info `yaml:"-"`
		Books []*Book    `yaml:"books,omitempty"`
	}

	// Book describes one book link loaded from the embedded YAML content.
	Book struct {
		Title string `yaml:"title,omitempty"`
		Link  string `yaml:"link,omitempty"`
	}
)
