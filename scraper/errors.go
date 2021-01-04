package scraper

import (
	"errors"
	"net/http"
)

// errInvalidURL is returned when a URL cannot be used to
// scrape data from the byte website
var errInvalidURL = errors.New("Invalid URL")

// IsErrInvalidURL checks if the given error is errInvalidURL
func IsErrInvalidURL(err error) bool {
	return err != nil && errors.Is(err, errInvalidURL)
}

// RequestError represents a non-200 response
type RequestError struct {
	StatusCode int
	Message    string
}

// NewRequestError creates a new RequestError instance
func NewRequestError(code int, message string) *RequestError {
	return &RequestError{code, message}
}

func (e *RequestError) Error() string {
	return e.Message
}

// IsStatusNotFound checks if the given error is a RequestError
// or wraps a RequestError and whether byte.co responded with
// a HTTP 404 status code
func IsStatusNotFound(err error) bool {
	var rerr *RequestError
	return err != nil && errors.As(err, &rerr) && rerr.StatusCode == http.StatusNotFound
}
