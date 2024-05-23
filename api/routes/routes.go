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
	return mux
}
