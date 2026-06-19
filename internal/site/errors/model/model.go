package model

import "github.com/alexfalkowski/web/internal/site/meta"

// NotFound is the view model rendered by the not-found page and fragment.
type NotFound struct {
	Info *meta.Info
	Page meta.Page
}
