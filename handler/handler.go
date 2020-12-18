package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjlock/byte-scraper-api/scraper"
)

type scraperHandler struct {
	scraper *scraper.Scraper
	router  *mux.Router
}

func NewScraperHandler() *scraperHandler {
	handler := &scraperHandler{
		scraper: scraper.NewScraper(),
		router:  mux.NewRouter(),
	}

	handler.routes()

	return handler
}

func (sh *scraperHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.router.ServeHTTP(w, r)
}

func (sh *scraperHandler) routes() {
	subrouter := sh.router.PathPrefix("/api").Methods("GET").Subrouter()

	// Routes
	subrouter.HandleFunc("/users/{username}", sh.getUser)
}
