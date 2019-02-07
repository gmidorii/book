package app

import (
	"github.com/gmidorii/book/server/domain"
	"github.com/gmidorii/book/server/form"
	"github.com/gmidorii/book/server/response"
)

type Book struct {
	br domain.BookRepository
}

func NewBook(br domain.BookRepository) Book {
	return Book{
		br: br,
	}
}

func (b Book) FetchList(form form.Books) (response.Books, error) {
	books := br.GetList(form.Offset, form.Limit)

	res := make([]response.Book, len(books))
	for i, b := range books {
		res[i] = response.Book{
			ID:       b.ID,
			Title:    b.Title,
			Author:   b.Author,
			Abstract: b.Abstract,
			ImgURL:   b.Img.URL,
		}
	}

	return response.Books{
		Items: res,
		Common: response.Common{
			Count:  len(res),
			Offset: form.Offset,
		},
	}, nil
}
