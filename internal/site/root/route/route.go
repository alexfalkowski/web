package route

import (
	"github.com/alexfalkowski/go-service/v3/net/http/mvc"
	"github.com/alexfalkowski/web/v2/internal/site/meta"
	"github.com/alexfalkowski/web/v2/internal/site/root/controller"
	"github.com/alexfalkowski/web/v2/internal/site/root/view"
)

// Register installs the GET and PUT root routes on the global MVC router.
func Register(server *mvc.Server, info *meta.Info) {
	view, partialView := view.NewRoot(server)

	server.Get("/{$}", controller.NewRoot(info, view), mvc.WithRouteUnauthenticated())
	server.Put("/{$}", controller.NewRoot(info, partialView), mvc.WithRouteUnauthenticated())
}
