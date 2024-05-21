package main

import (
	"log"
	"net/http"

	"github.com/sleepiinuts/webapp-plain/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting on port %s\n", port)
	http.ListenAndServe(port, nil)
}
