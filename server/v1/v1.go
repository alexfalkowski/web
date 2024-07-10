package v1

import (
	"embed"
)

//go:embed root.html
//go:embed home.html
//go:embed books.html
var fs embed.FS

func register() {
	rootRoute()
	homeRoute()
	booksRoute()
}
