package scraper

// RequestError represents a non-200 response
type RequestError struct {
	StatusCode int
	Message    string
}

func NewRequestError(code int, message string) *RequestError {
	return &RequestError{code, message}
}

func (e *RequestError) Error() string {
	return e.Message
}
