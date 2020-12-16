package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	errorResponse := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  code,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(&errorResponse)
	if err != nil {
		log.Println(err)
		return
	}
}

func respondInternalError(w http.ResponseWriter) {
	message := "Sorry, something went wrong on our side and we currently cannot handle the request."
	respondWithError(w, http.StatusInternalServerError, message)
}
