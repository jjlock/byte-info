package scraper

import (
	"strings"
)

type User struct {
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	Description     string `json:"description"`
	URL             string `json:"url"`
}

func (s *Scraper) ScrapeProfile(username string) (*User, error) {
	url := "https://byte.co/@" + username

	doc, err := s.scrape(url)
	if err != nil {
		return nil, err
	}

	user := &User{URL: url}

	sel := doc.Find(".author")

	user.Username = strings.TrimSpace(sel.Find(".username").Text())
	user.ProfileImageURL, _ = sel.Find(".avatar").Attr("src")
	user.Description = sel.Find(".bio").Text()

	return user, nil
}
