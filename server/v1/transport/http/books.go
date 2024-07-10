package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

type (
	// BookModel for HTTP.
	BookModel struct {
		Books []*Book
	}

	// Book for HTTP.
	Book struct {
		Title string
		Link  string
	}
)

// BookRoute registers routes for home.
func BooksRoute() {
	v := mvc.NewView(mvc.ParseTemplate(Views, "books.tmpl.html"), nil)

	mvc.Route("GET /books", v, func(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*BookModel, error) {
		m := &BookModel{
			Books: []*Book{
				{
					Title: "Kanban: Successful Evolutionary Change for Your Technology Business",
					Link:  "https://www.amazon.de/-/en/David-J-Anderson/dp/0984521402",
				},
			},
		}

		return m, nil
	})
}
