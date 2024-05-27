package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/handlers"
	"github.com/sleepiinuts/webapp-plain/internal/middlewares"
)

func Routes(h *handlers.Handler, ap *configs.AppProperties) http.Handler {
	mux := chi.NewRouter()

	// create middleware
	mw := middlewares.New(ap)

	mux.Use(mw.CSRFHandler)

	mux.Get("/", h.Home)
	mux.Get("/about", h.About)
	mux.Get("/contact", h.Contact)
	mux.Get("/rooms", h.Rooms)
	mux.Get("/make-reservation", h.MakeReservation)

	mux.Post("/make-reservation", h.PostMakeReservation)

	// set up static file FileSystem relative to main.go
	fs := http.FileServer(http.Dir("../web/static"))

	// if not stripPrefix, when serve /static, go will look for /web/static/static
	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	return mux
}
