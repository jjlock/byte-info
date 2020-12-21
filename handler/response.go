package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// respond is a helper that is responsible for sending a HTTP response
// in a request handler
func respond(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Failed to send response: ", err)
	}
}

// respondError is a helper that is responsible for sending an error response
// in a request handler with the given HTTP status code and message
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

// respondInternalError is a helper that can be used in place of respondError
// to send an error response with a default internal error code
func respondInternalError(w http.ResponseWriter) {
	message := "Sorry, something went wrong on our side and we currently cannot handle the request."
	respondError(w, http.StatusInternalServerError, message)
}
