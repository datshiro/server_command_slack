package consts

import "fmt"

type CustomError string

func (e CustomError) Error() string {
	return string(e)
}
func (e CustomError) Detail(err error) CustomError {
	return CustomError(fmt.Sprintf("%v : %v", e, err))
}

const (
	ErrorBadRequest CustomError = "Bad request"
)
