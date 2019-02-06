package handler

import (
	"io"
	"net/http"

	"github.com/gmidorii/book/server/form"
)

func BookList(w http.ResponseWriter, r *http.Request) {
	_, err := form.ParseBooks(r.Form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
}
