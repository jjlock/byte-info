package handler

import (
	"errors"
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
			message := "Unable to get user: " + requestError.Error()
			respondError(w, requestError.StatusCode, message)
		} else {
			respondInternalError(w)
		}
		return
	}

	respond(w, user, http.StatusOK)
}
