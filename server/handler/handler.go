package handler

import (
	"fmt"
	"io"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func errorMessage(w http.ResponseWriter, code int, message string) {
	w.Header().Add("content-type", "json")
	w.WriteHeader(code)
	io.WriteString(w, fmt.Sprintf("{\"message\": %v}", message))
}
