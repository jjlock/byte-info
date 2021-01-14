package scraper

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// roundTripFunc implements the http.RoundTripper interface and is the function
// used to return the mocked response.
type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// newTestScraper returns a Scraper instance for testing given the HTTP status
// code to return in the mocked response and the file for scraping.
func newTestScraper(t *testing.T, statusCode int, filename string) *Scraper {
	f := func(req *http.Request) *http.Response {
		if statusCode != 200 {
			return &http.Response{
				StatusCode: statusCode,
				Header:     make(http.Header),
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
		}

		html, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatal(err)
		}

		return &http.Response{
			StatusCode: statusCode,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(bytes.NewBuffer(html)),
		}
	}

	client := &http.Client{
		Transport: roundTripFunc(f),
		Timeout:   10 * time.Second,
	}

	return NewCustomScraper(client)
}
