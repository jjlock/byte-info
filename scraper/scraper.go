package scraper

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Scraper represents the scraper instance for scraping the byte website
type Scraper struct {
	client *http.Client
}

// NewScraper creates a new Scraper instance
func NewScraper() *Scraper {
	return &Scraper{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// get sends a GET request to the specifed url.
// A RequestError is returned on a non-200 response.
func (s *Scraper) get(url string) (*goquery.Document, error) {
	res, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err := errors.New("byte.co responded with HTTP status: " + res.Status)
		return nil, NewRequestError(res.StatusCode, err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return doc, nil
}
