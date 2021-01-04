package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jjlock/byte-scraper-api/scraper"
)

// respond sends a response with the given data and HTTP status code in JSON
func respond(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("handler: failed to send response: %v", err)
		return
	}
}

// handleError determines the appropriate error response to send based on the given error
func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		log.Println("handler: no error to send as an error response")
		return
	}

	var rerr *scraper.RequestError
	if errors.As(err, &rerr) {
		switch {
		case rerr.StatusCode >= 400 && rerr.StatusCode < 500:
			log.Printf("handler: %s", rerr.Error())
			respondInternalServerError(w)
		case rerr.StatusCode >= 500:
			respondError(w, http.StatusServiceUnavailable, "byte.co is currently unavailable")
		default:
			log.Printf("handler: byte.co responded with an unexpected HTTP status code: %d", rerr.StatusCode)
			respondInternalServerError(w)
		}
	} else {
		log.Printf("handler: %v", err)
		respondInternalServerError(w)
	}
}

// respondError sends an error response with the given HTTP status code and message in JSON
func respondError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  statusCode,
		Message: message,
	}

	respond(w, &errorResponse, statusCode)
}

// respondInternalError can be used in place of respondError to send an error response
// with a HTTP 500 status code
func respondInternalServerError(w http.ResponseWriter) {
	message := "Sorry, something went wrong on our side and we currently cannot handle the request."
	respondError(w, http.StatusInternalServerError, message)
}
