package validatorhelper

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/saidaydogan/chi-poc/pkg/errors"
)

func ToErrResponse(err error, translator ut.Translator) *errors.ErrResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := errors.ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}
		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s must be a maximum of %s in length", err.Field(), err.Param())
			case "url":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid URL", err.Field())
			default:
				resp.Errors[i] = err.Translate(translator)
			}
		}
		return &resp
	}
	return nil
}
