package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Failed to send response: ", err)
	}
}

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

func respondInternalError(w http.ResponseWriter) {
	message := "Sorry, something went wrong on our side and we currently cannot handle the request."
	respondError(w, http.StatusInternalServerError, message)
}
