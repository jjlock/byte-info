package scraper

import (
	"errors"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	client *http.Client
}

func NewScraper() *Scraper {
	return &Scraper{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *Scraper) scrape(url string) (*goquery.Document, error) {
	res, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err := errors.New("byte.co responded with HTTP status: " + res.Status)
		return nil, NewRequestError(res.StatusCode, err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
