package db

import "github.com/gmidorii/book/server/domain"

type book struct {
}

func NewBook() domain.BookRepository {
	return &book{}
}

func (b *book) Get(ID string) domain.Book {
	//TODO: fix
	return domain.Book{}
}

func (b *book) GetList(offset, limit int) []domain.Book {
	//TODO: fix
	return []domain.Book{domain.Book{ID: "1", Title: "book"}}
}
