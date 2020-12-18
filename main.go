package main

import (
	"log"
	"net/http"

	"github.com/jjlock/byte-scraper-api/handler"
)

func main() {
	srvHandler := handler.NewScraperHandler()

	log.Fatal(http.ListenAndServe(":8000", srvHandler))
}
