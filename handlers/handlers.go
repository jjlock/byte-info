package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-info/scraper"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := scraper.ScrapeProfile(vars["username"])
	if err != nil {
		var requestError *scraper.RequestError
		if errors.As(err, &requestError) {
			respondWithError(w, requestError.StatusCode, requestError.Error())
		} else {
			respondInternalError(w)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		return
	}
}
