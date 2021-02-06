package errors

import "fmt"

type ErrResponse struct {
	Errors []string `json:"errors"`
}

type CustomError interface {
	Error() string
	Code() string
	Message() string
}

type CustomErrorDef struct {
	code    string
	message string
	params  []interface{}
}

func Error(code, message string) CustomErrorDef {
	return CustomErrorDef{
		code:    code,
		message: message,
	}
}

func Errorf(code, message string, params []interface{}) CustomErrorDef {
	return CustomErrorDef{
		code:    code,
		message: message,
		params:  params,
	}
}

func (e CustomErrorDef) Error() string {
	return e.message
}

func (e CustomErrorDef) Code() string {
	return e.code
}

func (e CustomErrorDef) Message() string {
	return fmt.Sprintf(e.message, e.params)
}

func (e CustomErrorDef) Equal(err error) bool {
	switch err.(type) {
	case CustomError:
		var ourError = err.(CustomError)

		return ourError.Code() == e.code
	default:
		return false
	}
}
