package scraper

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

// get sends a GET request to the specifed url and returns a goquery.Document.
// A RequestError is returned on a non-200 response.
func (s *Scraper) get(url string) (*goquery.Document, error) {
	if !isByteURL(url) {
		fmt.Println("invalid url")
		return nil, errors.New("Invalid URL")
	}

	res, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, NewRequestError(res.StatusCode, "byte.co responded with HTTP status: "+res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return doc, nil
}

func isByteURL(reqURL string) bool {
	u, err := url.ParseRequestURI(reqURL)
	return err == nil && u.Scheme == "https" && u.Host == "byte.co"
}
