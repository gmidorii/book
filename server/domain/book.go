package domain

type Book struct {
	ID       string
	Title    string
	Author   string
	Abstract string
	Img      Img
}

type Img struct {
	Title string
	URL   string
}

type BookRepository interface {
	Get(ID string) Book
	GetList(offset, limit int) []Book
}
