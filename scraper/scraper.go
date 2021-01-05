package scraper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Scraper represents the scraper instance for scraping the byte website
type Scraper struct {
	baseURL string
	client  *http.Client
}

// NewScraper creates a new Scraper instance
func NewScraper() *Scraper {
	return &Scraper{
		baseURL: "https://byte.co",
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

// get sends a GET request to the specifed url.
// A RequestError is returned on a non-200 response, otherwise it returns
// any error returned from sending the request or parsing the response.
func (s *Scraper) get(url string) (*goquery.Document, error) {
	res, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, &RequestError{
			StatusCode: res.StatusCode,
			Message:    "byte.co responded with HTTP status: " + res.Status,
		}
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read response: %v", err)
	}

	return doc, nil
}
