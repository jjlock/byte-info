package scraper

type RequestError struct {
	StatusCode int
	Err        error
}

func NewRequestError(code int, err error) *RequestError {
	return &RequestError{code, err}
}

func (e *RequestError) Error() string {
	return e.Err.Error()
}
