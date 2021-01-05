package scraper

import (
	"strings"
)

// User represents a byte user
type User struct {
	Username        string   `json:"username"`
	ProfileImageURL string   `json:"profile_image_url"`
	Description     string   `json:"description"`
	RecentByteURLs  []string `json:"recent_byte_urls"`
	URL             string   `json:"url"`
}

// GetUser returns scraped user data given a username.
// A RequestError is returned on a non-200 response, otherwise it returns
// any error returned from sending the request or parsing the response.
func (s *Scraper) GetUser(username string) (*User, error) {
	url := s.baseURL + "/@" + username

	doc, err := s.get(url)
	if err != nil {
		return nil, err
	}

	user := &User{URL: url}

	sel := doc.Find(".author")

	user.Username = strings.TrimSpace(sel.Find(".username").Text())
	user.ProfileImageURL, _ = sel.Find(".avatar").Attr("src")
	user.Description = sel.Find(".bio").Text()
	user.RecentByteURLs = make([]string, 0)

	sel = doc.Find(".post")
	for i := 0; i < len(sel.Nodes); i++ {
		single := sel.Eq(i)
		href, _ := single.Find("a").Attr("href")
		user.RecentByteURLs = append(user.RecentByteURLs, s.baseURL+href)
	}

	return user, nil
}
