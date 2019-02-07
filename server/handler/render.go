package handler

import (
	"encoding/json"
	"errors"
	"io"
)

func renderJSON(w io.Writer, i interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func renderHTML() error {
	return errors.New("unsupport")
}
