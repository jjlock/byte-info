package scraper

import (
	"fmt"
	"strconv"
	"strings"
)

// Byte represents a post (called a byte)
type Byte struct {
	ID           string   `json:"id"`
	User         string   `json:"user"`
	UserURL      string   `json:"user_url"`
	ThumbnailURL string   `json:"thumbnail_url"`
	Caption      string   `json:"caption"`
	CreatedAt    string   `json:"created_at"`
	Loops        int      `json:"loops"`
	URLs         []string `json:"urls"`
}

// GetByte returns scraped data of a byte given its ID.
// A RequestError is returned on a non-200 response, otherwise it returns
// any error returned from sending the request or parsing the response.
func (s *Scraper) GetByte(id string) (*Byte, error) {
	url := ByteBaseURL + "/b/" + id
	doc, err := s.get(url)
	if err != nil {
		return nil, err
	}

	byte := &Byte{ID: id}

	byte.ThumbnailURL, _ = doc.Find(`#vinit`).Attr("poster")

	sel := doc.Find("#desktop div:not([class])")

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

	byte.URLs = []string{byte.UserURL + "/" + id, url}

	return byte, nil
}
