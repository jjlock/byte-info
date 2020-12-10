package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-info/scraper"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := scraper.ScrapeProfile(vars["username"])
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
