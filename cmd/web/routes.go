package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", h.Home)
	mux.Get("/about", h.About)
	return mux
}
