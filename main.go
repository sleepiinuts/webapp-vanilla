package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting on port: \"%s\"\n", port)
	http.ListenAndServe(port, nil)
}
