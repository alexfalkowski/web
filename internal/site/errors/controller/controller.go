package controller

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/meta"
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/errors/model"
	site "github.com/alexfalkowski/web/internal/site/meta"
)

// NewNotFound returns an MVC not-found controller that renders the full page for
// normal requests and the partial fragment for PUT requests.
func NewNotFound(info *site.Info, view *mvc.View, partialView *mvc.View) mvc.NotFoundController[model.NotFound] {
	return func(ctx context.Context) (*mvc.View, *model.NotFound) {
		notFound := &model.NotFound{
			Info: info,
			Page: site.Page{
				Title:       "Page not found | Lean Thoughts",
				Description: "The requested Lean Thoughts page could not be found.",
			},
		}
		req := meta.Request(ctx)

		switch req.Method {
		case http.MethodPut:
			return partialView, notFound
		default:
			return view, notFound
		}
	}
}
