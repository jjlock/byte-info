package scraper

import (
	"errors"
	"net/http"
)

// RequestError should be returned when byte.co responses with
// a non-200 HTTP status code.
type RequestError struct {
	StatusCode int
	Message    string
}

func (e *RequestError) Error() string {
	return e.Message
}

// IsStatusNotFound checks if the given error is RequestError
// or wraps RequestError and whether byte.co responded with
// a 404 HTTP status code.
func IsStatusNotFound(err error) bool {
	var rerr *RequestError
	return err != nil && errors.As(err, &rerr) && rerr.StatusCode == http.StatusNotFound
}
