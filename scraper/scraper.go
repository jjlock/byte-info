package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type User struct {
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	Description     string `json:"description"`
	URL             string `json:"url"`
}

func ScrapeProfile(username string) (*User, error) {
	url := "https://byte.co/@" + username
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err := fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
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
