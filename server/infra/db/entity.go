package db

type BookEntity struct {
	ISBN     string
	Title    string
	Author   string
	Abstract string
}

type ImgEntity struct {
	ID    int
	ISBN  string
	Title string
	URL   string
}
