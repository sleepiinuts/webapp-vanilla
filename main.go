package main

import (
	"fmt"
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

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the home page")
	if err != nil {
		log.Fatalf("response writing error: %v", err)
	}

	log.Printf("Number of byte written: %d\n", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the about page")
	if err != nil {
		log.Fatalf("response writing error: %v", err)
	}

	log.Printf("Number of byte written: %d\n", n)
}
