package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-scraper-api/scraper"
)

// getUser gets a user by their username
func (sh *ScraperHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := sh.scraper.GetUser(vars["username"])
	if err != nil {
		if scraper.IsStatusNotFound(err) {
			message := "User not found. User either does not exist or does exist but has no posts."
			respondError(w, http.StatusNotFound, message)
		} else {
			handleError(w, err)
		}
		return
	}

	respond(w, user, http.StatusOK)
}

// getByte gets a byte by a url
func (sh *ScraperHandler) getByte(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	byte, err := sh.scraper.GetByte(url)
	if err != nil {
		switch {
		case scraper.IsErrInvalidURL(err):
			respondError(w, http.StatusBadRequest, "Invalid URL. The URL must link to a byte.")
		case scraper.IsStatusNotFound(err):
			respondError(w, http.StatusNotFound, "Byte not found")
		default:
			handleError(w, err)
		}
		return
	}

	respond(w, byte, http.StatusOK)
}
