package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sleepiinuts/webapp-plain/api/handlers"
	"github.com/sleepiinuts/webapp-plain/api/middlewares"
	"github.com/sleepiinuts/webapp-plain/configs"
)

func Routes(h *handlers.Handler, ap *configs.AppProperties) http.Handler {
	mux := chi.NewRouter()

	// create middleware
	mw := middlewares.New(ap)

	mux.Use(mw.CSRFHandler)

	mux.Get("/", h.Home)
	mux.Get("/about", h.About)

	// set up static file FileSystem relative to main.go
	fs := http.FileServer(http.Dir("../web/static"))

	// if not stripPrefix, when serve /static, go will look for /web/static/static
	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	return mux
}
