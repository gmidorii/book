package app

import "github.com/gmidorii/book/server/domain"

type Book struct {
	br domain.BookRepository
}

func NewBook(br domain.BookRepository) Book {
	return Book{
		br: br,
	}
}

func (b Book) FetchList() []domain.Book {
}
