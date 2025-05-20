package route

import (
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/meta"
	"github.com/alexfalkowski/web/internal/site/root/controller"
	"github.com/alexfalkowski/web/internal/site/root/view"
)

// Register root.
func Register(info *meta.Info) {
	view, partialView := view.NewRoot()

	mvc.Get("/", controller.NewRoot(info, view))
	mvc.Put("/", controller.NewRoot(info, partialView))
}
