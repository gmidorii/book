package handler

import (
	"log"
	"net/http"

	"github.com/gmidorii/book/server/form"
	"github.com/pkg/errors"
)

func Books(w http.ResponseWriter, r *http.Request) {
	_, err := form.ParseBooks(r.URL.Query())
	if err != nil {
		log.Println(err)
		errorMessage(w, http.StatusBadRequest, errors.Cause(err).Error())
		return
	}
}
