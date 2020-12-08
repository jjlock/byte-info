package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

type User struct {
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	Description     string `json:"description"`
	URL             string `json:"url"`
}

func scrapeProfile(username string) (*User, error) {
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

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := scrapeProfile(vars["username"])
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Methods("GET").Subrouter()

	// Routes
	s.HandleFunc("/users/{username}", getUser)

	log.Fatal(http.ListenAndServe(":8000", r))
}
