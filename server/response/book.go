package response

type Books struct {
	Items  []Book `json:"items"`
	Common `json:"common"`
}

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Abstract string `json:"abstract,omit_empty"`
	ImgURL   string `json:"img_url,omit_empty"`
}
