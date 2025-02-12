package statuserr

import "net/http"

type NotFoundErr struct {
	message    string
	statusCode int
}

func NewNotFoundErr(message string) *NotFoundErr {
	return &NotFoundErr{
		message:    message,
		statusCode: http.StatusNotFound,
	}
}

func (b *NotFoundErr) StatusCode() int {
	return b.statusCode
}

func (b *NotFoundErr) Error() string {
	return b.message
}
