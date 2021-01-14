package scraper

import (
	"strings"
)

// User represents a byte user.
type User struct {
	Username        string   `json:"username"`
	ProfileImageURL string   `json:"profile_image_url"`
	Bio             string   `json:"bio"`
	RecentByteIDs   []string `json:"recent_byte_ids"`
	RecentByteURLs  []string `json:"recent_byte_urls"`
	URL             string   `json:"url"`
}

// GetUser returns scraped user data given a username.
// *RequestError is returned on a non-200 response, otherwise it returns
// any error returned from sending the request or parsing the response.
func (s *Scraper) GetUser(username string) (*User, error) {
	url := ByteBaseURL + "/@" + username
	doc, err := s.get(url)
	if err != nil {
		return nil, err
	}

	user := &User{URL: url}
	sel := doc.Find(".author")

	user.Username = strings.TrimSpace(sel.Find(".username").Text())
	user.ProfileImageURL, _ = sel.Find(".avatar").Attr("src")
	user.Bio = sel.Find(".bio").Text()
	user.RecentByteIDs = []string{}
	user.RecentByteURLs = []string{}

	sel = doc.Find(".post")
	for i := 0; i < len(sel.Nodes); i++ {
		single := sel.Eq(i)
		href, _ := single.Find("a").Attr("href")
		user.RecentByteIDs = append(user.RecentByteIDs, strings.TrimPrefix(href, "/@"+user.Username+"/"))
		user.RecentByteURLs = append(user.RecentByteURLs, ByteBaseURL+href)
	}

	return user, nil
}
