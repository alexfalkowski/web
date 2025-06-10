package books

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/web/internal/site/books/repository"
	"github.com/alexfalkowski/web/internal/site/books/route"
)

// Module for fx.
var Module = di.Module(
	repository.Module,
	route.Module,
)
