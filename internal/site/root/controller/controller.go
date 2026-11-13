package controller

import (
	"github.com/alexfalkowski/go-service/v3/context"
	"github.com/alexfalkowski/go-service/v3/net/http/mvc"
	"github.com/alexfalkowski/web/v2/internal/site/meta"
	"github.com/alexfalkowski/web/v2/internal/site/root/model"
)

// NewRoot returns an MVC controller that renders the root model with shared site
// metadata.
func NewRoot(info *meta.Info, view *mvc.View) mvc.Controller[model.Root] {
	return func(_ context.Context) (*mvc.View, *model.Root, error) {
		root := &model.Root{
			Info: info,
			Page: meta.Page{
				Title:       "root.title",
				Description: "root.description",
			},
		}

		return view, root, nil
	}
}
