package handler

import (
	"log"
	"net/http"

	"github.com/gmidorii/book/server/form"
	"github.com/gmidorii/book/server/provider"
	"github.com/pkg/errors"
)

func Books(w http.ResponseWriter, r *http.Request) {
	f, err := form.ParseBooks(r.URL.Query())
	if err != nil {
		log.Println(err)
		errorMessage(w, http.StatusBadRequest, errors.Cause(err).Error())
		return
	}

	app, err := provider.InitBookApp()
	if err != nil {
		log.Println(err)
		errorMessage(w, http.StatusInternalServerError, "Server Error")
		return
	}

	res, err := app.FetchList(f)
	if err != nil {
		log.Println(err)
		errorMessage(w, http.StatusInternalServerError, "Server Error")
		return
	}

	renderJSON(w, &res)
}
