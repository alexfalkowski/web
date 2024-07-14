package site

import (
	"embed"
)

//go:embed **/*.html
//go:embed **/*.yaml
var fs embed.FS
