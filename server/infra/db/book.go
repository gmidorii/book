package db

import (
	"database/sql"

	"github.com/gmidorii/book/server/domain"
	"github.com/jmoiron/sqlx"
)

type book struct {
	db *sql.DB
}

func NewBook(db *sql.DB) domain.BookRepository {
	return &book{db: db}
}

func (b *book) Get(ID string) (domain.Book, error) {
	//TODO: fix
	return domain.Book{}, nil
}

const (
	getListColumn = `
select
  books.isbn,
  books.title,
  books.abstract,
  author.name as author_name,
  img.title as img_title,
  img.url as img_url
from
 (
    select
      isbn, title, abstract
    from book.book_t
    LIMIT ?,?
) as books
inner join 
book.author_t as author
on
books.isbn = author.isbn
inner join
book.img_t as img
on
books.isbn = img.isbn
`
)

type GetListColumn struct {
	ISBN     string
	Title    string
	Abstract string
	Author   string `db:"author_name"`
	ImgTitle string `db:"img_title"`
	ImgURL   string `db:"img_url"`
}

func (b *book) GetList(offset, limit int) ([]domain.Book, error) {
	db := sqlx.NewDb(b.db, "mysql")
	e := []GetListColumn{}
	if err := db.Select(&e, getListColumn, offset, limit); err != nil {
		return nil, err
	}
	return []domain.Book{domain.Book{ID: "1", Title: "book"}}, nil
}
