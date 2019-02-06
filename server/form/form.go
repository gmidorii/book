package form

import (
	"fmt"

	"github.com/pkg/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

var v *validator.Validate = validator.New()

func validate(i interface{}) error {
	if err := v.Struct(i); err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.Wrap(err, "invalidation error")
		}

		var s string
		for _, err := range err.(validator.ValidationErrors) {
			s = fmt.Sprintf("%v tag: %v, field: %v, param: %v | ", s, err.Tag(), err.Field(), err.Param())
		}

		return errors.Wrap(err, s)
	}
	return nil
}
