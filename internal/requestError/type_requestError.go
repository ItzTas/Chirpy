package requesterror

type RequestError struct {
	StatusCode int
	message    string
}

func (e *RequestError) Error() string {
	return e.message
}

func (e *RequestError) Status() int {
	return e.StatusCode
}

func NewRequestErr(statusCode int, message string) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		message:    message,
	}
}
