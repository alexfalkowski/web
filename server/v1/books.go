package v1

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

func booksRoute() {
	mvc.Route(
		"GET /books",
		mvc.NewSuccessView(mvc.ParseTemplate(fs, "books.html")),
		func(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*BookModel, error) {
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
