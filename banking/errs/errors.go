package errs

import "net/http"

type AppErr struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusNotFound,
	}
}
func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
