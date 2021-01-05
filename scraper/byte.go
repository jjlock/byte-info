package scraper

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Byte represents a post (called a byte)
type Byte struct {
	User      string `json:"user"`
	UserURL   string `json:"user_url"`
	Caption   string `json:"caption"`
	CreatedAt string `json:"created_at"`
	Loops     int    `json:"loops"`
	URL       string `json:"url"`
}

// GetByte returns scraped data on a byte given a URL to the byte.
// An InvalidURLError is returned if the given URL is not a link to a byte.
// A RequestError is returned on a non-200 response, otherwise it returns
// any error returned from sending the request or parsing the response.
func (s *Scraper) GetByte(url string) (*Byte, error) {
	if !s.isValidURL(url) {
		return nil, &InvalidURLError{Reason: url + " is not a link to a byte"}
	}

	doc, err := s.get(url)
	if err != nil {
		return nil, err
	}

	sel := doc.Find("#desktop div:not([class])")

	byte := &Byte{URL: url}

	byte.User = sel.Find(".username a").Text()

	href, _ := sel.Find(".username a").Attr("href")
	byte.UserURL = ByteBaseURL + href

	byte.Caption = sel.Find(".post-content").Text()
	byte.CreatedAt = sel.Find(".avatar-wrapper div:not([class])").Text()

	loopsText := strings.TrimSpace(sel.Find(".loops").Text())
	loops := strings.ReplaceAll(loopsText, ",", "")
	byte.Loops, err = strconv.Atoi(loops)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse response: %v", err)
	}

	return byte, nil
}

// isValidURL checks if the given URL is a link to a byte
func (s *Scraper) isValidURL(rawurl string) bool {
	ubase, err := url.ParseRequestURI(ByteBaseURL)
	if err != nil {
		return false
	}
	u, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	matched, err := regexp.MatchString(`^/@?[0-9A-Za-z]+/[0-9A-Za-z]+$`, u.Path)
	if err != nil {
		return false
	}

	return ubase.Scheme == u.Scheme && ubase.Host == u.Host && matched
}
