package handler

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

// getUser gets a user by their username
func (sh *ScraperHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := sh.scraper.GetUser(vars["username"])
	if err != nil {
		if isErrNotFound(err) {
			message := "User not found. User either does not exist or does exist but has not made a post."
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
	matched, err := regexp.MatchString(`^https://byte\.co/.+/.+$`, url)
	if err != nil {
		log.Printf("handler: %v", err)
		respondInternalServerError(w)
		return
	}
	if !matched {
		respondError(w, http.StatusBadRequest, "Invalid URL. The URL must link to a byte.")
		return
	}

	byte, err := sh.scraper.GetByte(url)
	if err != nil {
		if isErrNotFound(err) {
			respondError(w, http.StatusNotFound, "Byte not found")
		} else {
			handleError(w, err)
		}
		return
	}

	respond(w, byte, http.StatusOK)
}
