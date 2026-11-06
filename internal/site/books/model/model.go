package model

import "github.com/alexfalkowski/web/internal/site/meta"

// Books is the view model rendered by the books page.
type Books struct {
	Info  *meta.Info `yaml:"-"`
	Page  meta.Page  `yaml:"-"`
	Books []*Book    `yaml:"books,omitempty"`
}

// Book describes one book recommendation loaded from the embedded YAML content.
// Topics and Note are catalog message IDs resolved by the request's MVC localizer;
// titles, authors, and links retain their original values in every language.
type Book struct {
	Title  string `yaml:"title,omitempty"`
	Author string `yaml:"author,omitempty"`
	Topics string `yaml:"topics,omitempty"`
	Note   string `yaml:"note,omitempty"`
	Link   string `yaml:"link,omitempty"`
}
