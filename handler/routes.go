package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-scraper-api/scraper"
)

// getUser gets a user by their username
func (sh *ScraperHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := sh.scraper.GetUser(vars["username"])
	if err != nil {
		var requestError *scraper.RequestError
		if errors.As(err, &requestError) {
			message := "Unable to get user: " + requestError.Error()
			respondError(w, requestError.StatusCode, message)
		} else {
			respondInternalError(w)
		}
		return
	}

	respond(w, user, http.StatusOK)
}

func (sh *ScraperHandler) getByte(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	byte, err := sh.scraper.GetByte(url)
	if err != nil {
		var requestError *scraper.RequestError
		if errors.As(err, &requestError) {
			message := "Unable to get byte: " + requestError.Error()
			respondError(w, requestError.StatusCode, message)
		} else {
			respondInternalError(w)
		}
		return
	}

	respond(w, byte, http.StatusOK)
}
