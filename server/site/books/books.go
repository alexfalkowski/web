package books

import (
	"cmp"
	"context"
	"io/fs"
	"net/http"
	"slices"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

type (
	// BookModel for v1.
	BookModel struct {
		Books []*Book
	}

	// Book for v1.
	Book struct {
		Title string
		Link  string
	}
)

// Path for books.
func Path() string {
	return "GET /books"
}

// View for books.
func View(fs fs.FS) *mvc.View {
	return mvc.NewSuccessView(mvc.ParseTemplate(fs, "books/success.html"))
}

// Controller for books.
func Controller(_ context.Context, _ *http.Request, _ http.ResponseWriter) (*BookModel, error) {
	books := []*Book{
		{
			Title: "Kanban: Successful Evolutionary Change for Your Technology Business",
			Link:  "https://www.amazon.de/-/en/David-J-Anderson/dp/0984521402",
		},
		{
			Title: "Team Topologies: Organizing Business and Technology Teams for Fast Flow",
			Link:  "https://www.amazon.de/-/en/Team-Topologies-Organizing-Business-Technology/dp/1942788819",
		},
		{
			Title: "Modern Software Engineering: Doing What Works to Build Better Software Faster",
			Link:  "https://www.amazon.de/-/en/Modern-Software-Engineering-Better-Faster/dp/0137314914",
		},
	}

	slices.SortFunc(books, func(a, b *Book) int {
		return cmp.Compare(a.Title, b.Title)
	})

	m := &BookModel{Books: books}

	return m, nil
}
