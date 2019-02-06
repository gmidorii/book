package form

import (
	"net/url"

	"github.com/pkg/errors"
)

type Books struct {
	Offset int `validate:"min=0"`
	Limit  int `validate:"min=0,max=100"`
}

func ParseBooks(values url.Values) (Books, error) {
	b := Books{}
	if err := validate(b); err != nil {
		return Books{}, errors.Wrap(err, "parse books error")
	}

	return b, nil
}
