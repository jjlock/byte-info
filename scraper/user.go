package scraper

import (
	"fmt"
	"strings"
)

// User represents a byte user
type User struct {
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	Description     string `json:"description"`
	URL             string `json:"url"`
}

// GetUser returns scraped user data given a username
func (s *Scraper) GetUser(username string) (*User, error) {
	url := "https://byte.co/@" + username

	doc, err := s.get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to get user: %w", err)
	}

	user := &User{URL: url}

	sel := doc.Find(".author")

	user.Username = strings.TrimSpace(sel.Find(".username").Text())
	user.ProfileImageURL, _ = sel.Find(".avatar").Attr("src")
	user.Description = sel.Find(".bio").Text()

	return user, nil
}
