package route

import (
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/web/internal/site/errors/controller"
	"github.com/alexfalkowski/web/internal/site/errors/view"
	"github.com/alexfalkowski/web/internal/site/meta"
)

// Register installs the global not-found handler on the MVC router.
func Register(info *meta.Info) {
	view, partialView := view.NewNotFound()

	mvc.NotFound(controller.NewNotFound(info, view, partialView))
}
