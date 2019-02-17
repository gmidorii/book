package db

type BookEntity struct {
	ISBN     string
	Title    string
	Abstract string
}

type AuthorEntity struct {
	ID   int
	Name string
	ISBN string
}

type ImgEntity struct {
	ID    int
	ISBN  string
	Title string
	URL   string
}
