package statuserr

type ApiError interface {
	StatusCode() int
	Error() string
}
