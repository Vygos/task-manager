package statuserr

import "net/http"

type BadRequestErr struct {
	message    string
	statusCode int
}

func NewBadRequest(message string) *BadRequestErr {
	return &BadRequestErr{
		message:    message,
		statusCode: http.StatusBadRequest,
	}
}

func (b *BadRequestErr) StatusCode() int {
	return b.statusCode
}

func (b *BadRequestErr) Error() string {
	return b.message
}
