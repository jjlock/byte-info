package scraper

import (
	"errors"
	"net/http"
)

// RequestError should be returned when byte.co responses with
// a non-200 HTTP status code
type RequestError struct {
	StatusCode int
	Message    string
}

func (e *RequestError) Error() string {
	return e.Message
}

// IsStatusNotFound checks if the given error is a RequestError
// or wraps a RequestError and whether byte.co responded with
// a 404 HTTP status code
func IsStatusNotFound(err error) bool {
	var rerr *RequestError
	return err != nil && errors.As(err, &rerr) && rerr.StatusCode == http.StatusNotFound
}

// InvalidURLError should be returned when a URL cannot be used
// to scrape data from the byte website
type InvalidURLError struct {
	Reason string
}

func (e *InvalidURLError) Error() string {
	return "Invalid URL: " + e.Reason
}

// IsErrInvalidURL checks if the given error is an InvalidURLError
// or wraps an InvalidURLError
func IsErrInvalidURL(err error) bool {
	var iuerr *InvalidURLError
	return err != nil && errors.As(err, &iuerr)
}
