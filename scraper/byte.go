package scraper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Byte represents a post (called a byte)
type Byte struct {
	User      string `json:"user"`
	Caption   string `json:"caption"`
	CreatedAt string `json:"created_at"`
	Loops     int    `json:"loops"`
}

// GetByte returns scraped data on a byte given a url to the byte
func (s *Scraper) GetByte(url string) (*Byte, error) {
	doc, err := s.get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to get byte: %w", err)
	}

	if sel := doc.Find("#post"); !sel.Is("#post") {
		return nil, errors.New("The URL is not a link to a byte")
	}

	sel := doc.Find("#desktop div:not([class])")

	byte := &Byte{}

	byte.User = sel.Find(".username a").Text()
	byte.Caption = sel.Find(".post-content").Text()
	byte.CreatedAt = sel.Find(".avatar-wrapper div:not([class])").Text()

	loopsText := strings.TrimSpace(sel.Find(".loops").Text())
	loops := strings.ReplaceAll(loopsText, ",", "")
	byte.Loops, err = strconv.Atoi(loops)
	if err != nil {
		return nil, fmt.Errorf("Could not convert loops text to int: %v", err)
	}

	return byte, nil
}
