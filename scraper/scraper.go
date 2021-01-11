// Package scraper implements a framework for scraping the byte website
package scraper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// ByteBaseURL is the base URL of the byte website.
// It is used by a Scraper to construct URLs for scraping data.
const ByteBaseURL = "https://byte.co"

// ByteScraper implements the Scraper interface and represents the
// instance used for scraping the byte website
type ByteScraper struct {
	client *http.Client
}

// NewByteScraper creates a new ByteScraper instance
func NewByteScraper() *ByteScraper {
	return &ByteScraper{client: &http.Client{Timeout: 10 * time.Second}}
}

// NewCustomByteScraper creates a new ByteScraper instance with the given http.Client
func NewCustomByteScraper(c *http.Client) *ByteScraper {
	return &ByteScraper{client: c}
}

// get sends a GET request to the specifed url.
// A RequestError is returned on a non-200 response, otherwise it returns
// any error returned from sending the request or parsing the response.
func (s *ByteScraper) get(url string) (*goquery.Document, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &RequestError{
			StatusCode: resp.StatusCode,
			Message:    "byte.co responded with HTTP status: " + resp.Status,
		}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read response: %v", err)
	}

	return doc, nil
}
