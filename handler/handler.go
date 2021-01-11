// Package handler provides a http.Handler that handles requests to scrape the byte website.
package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-scraper-api/scraper"
)

// Scraper represents the ability to scrape the byte website.
type Scraper interface {
	GetUser(username string) (*scraper.User, error)
	GetByte(id string) (*scraper.Byte, error)
}

// ScraperHandler implements the http.Handler interface and
// handles request to scrape the byte website.
type ScraperHandler struct {
	scraper Scraper
	router  *mux.Router
}

// NewScraperHandler creates a new ScraperHandler instance with
// all the routes registered to the router.
func NewScraperHandler() *ScraperHandler {
	handler := &ScraperHandler{
		scraper: scraper.NewScraper(),
		router:  mux.NewRouter(),
	}

	handler.routes()

	return handler
}

// ServeHTTP calls the router's ServeHTTP method.
func (sh *ScraperHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.router.ServeHTTP(w, r)
}

// routes registers the routes to the router.
func (sh *ScraperHandler) routes() {
	subrouter := sh.router.PathPrefix("/api").Methods("GET").Subrouter()

	// Routes
	subrouter.HandleFunc("/users/{username}", sh.getUser)
	subrouter.HandleFunc("/bytes/{id}", sh.getByte)
}
