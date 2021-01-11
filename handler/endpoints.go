package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-scraper-api/scraper"
)

// getUser handles getting a user by their username.
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

	respond(w, http.StatusOK, user)
}

// getByte handles getting a byte by its ID.
func (sh *ScraperHandler) getByte(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	byte, err := sh.scraper.GetByte(vars["id"])
	if err != nil {
		if scraper.IsStatusNotFound(err) {
			respondError(w, http.StatusNotFound, "Byte not found.")
		} else {
			handleError(w, err)
		}
		return
	}

	respond(w, http.StatusOK, byte)
}
