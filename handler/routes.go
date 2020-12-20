package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-scraper-api/scraper"
)

func (sh *scraperHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := sh.scraper.GetUser(vars["username"])
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
