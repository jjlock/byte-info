package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-info/handlers"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Methods("GET").Subrouter()

	// Routes
	s.HandleFunc("/users/{username}", handlers.GetUser)

	log.Fatal(http.ListenAndServe(":8000", r))
}
