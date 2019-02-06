package form

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

const (
	dOffset = 0
	dLimit  = 10
)

type Books struct {
	Offset int `validate:"min=0"`
	Limit  int `validate:"min=0,max=100"`
}

func ParseBooks(values url.Values) (Books, error) {
	b := Books{
		Offset: dOffset,
		Limit:  dLimit,
	}
	var err error

	offset := values.Get("offset")
	if offset != "" {
		b.Offset, err = strconv.Atoi(offset)
		if err != nil {
			return Books{}, errors.New("failed type offset (int)")
		}
	}

	limit := values.Get("limit")
	if limit != "" {
		b.Limit, err = strconv.Atoi(offset)
		if err != nil {
			return Books{}, errors.New("failed type limit (int)")
		}
	}

	if err := validate(b); err != nil {
		return Books{}, errors.Wrap(err, "parse books error")
	}

	return b, nil
}
