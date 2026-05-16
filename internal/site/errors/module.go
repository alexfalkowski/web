package errors

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/web/internal/site/errors/route"
)

// Module for fx.
var Module = di.Module(
	route.Module,
)
