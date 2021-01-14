package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jjlock/byte-scraper-api/scraper"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// respond sends a response with the given data and HTTP status code in JSON.
func respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Failed to send response: %v", err)
		return
	}
}

// handleError determines the appropriate error response to send based on the given error.
func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		log.Println("No error to send as an error response")
		return
	}

	var rerr *scraper.RequestError
	if errors.As(err, &rerr) {
		switch {
		case rerr.StatusCode >= 400 && rerr.StatusCode < 500:
			log.Println("byte.co responded with HTTP status: " + http.StatusText(rerr.StatusCode))
			respondInternalServerError(w)
		case rerr.StatusCode >= 500:
			respondError(w, 503, "byte.co is currently unavailable.")
		default:
			log.Printf("byte.co responded with an unexpected HTTP status code: %d", rerr.StatusCode)
			respondInternalServerError(w)
		}
		return
	}

	log.Printf("Could not handle response from byte.co: %v", err)
	respondInternalServerError(w)
}

// respondError sends an error response with the given HTTP status code and message in JSON.
func respondError(w http.ResponseWriter, statusCode int, message string) {
	respond(w, statusCode, errorResponse{Status: statusCode, Message: message})
}

// respondInternalError can be used in place of respondError to send an error response with
// a 500 HTTP status code and default message in JSON.
func respondInternalServerError(w http.ResponseWriter) {
	respond(w, 500, errorResponse{Status: 500, Message: "Sorry, something went wrong on our side and we currently cannot handle the request."})
}
