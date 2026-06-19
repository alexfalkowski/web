package model

import "github.com/alexfalkowski/web/internal/site/meta"

// Root is the view model rendered by the home page and fragment.
type Root struct {
	Info *meta.Info
	Page meta.Page
}
