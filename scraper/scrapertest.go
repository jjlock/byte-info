package scraper

import (
	"net/http"
	"time"
)

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestScraper creates a new scraper.Scraper instance to use for testing
func NewTestScraper(f roundTripFunc) *Scraper {
	ts := NewScraper()
	client := &http.Client{
		Transport: roundTripFunc(f),
		Timeout:   10 * time.Second,
	}
	ts.SetClient(client)

	return ts
}
