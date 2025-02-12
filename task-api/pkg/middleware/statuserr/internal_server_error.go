package statuserr

import "net/http"

type InternalServerError struct {
	message    string
	statusCode int
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{
		message:    message,
		statusCode: http.StatusInternalServerError,
	}
}

func (i *InternalServerError) StatusCode() int {
	return i.statusCode
}

func (i *InternalServerError) Error() string {
	return i.message
}
