package books

import (
	"github.com/alexfalkowski/web/internal/site/books/repository"
	"github.com/alexfalkowski/web/internal/site/books/route"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	repository.Module,
	route.Module,
)
