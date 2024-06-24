package routes

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/handlers"
	"github.com/sleepiinuts/webapp-plain/internal/middlewares"
)

func Routes(h *handlers.Handler, ap *configs.AppProperties, sm *scs.SessionManager) http.Handler {
	mux := chi.NewRouter()

	// create middleware
	mw := middlewares.New(ap, sm)

	mux.Use(mw.CSRFHandler)

	mux.Get("/", h.Home)
	mux.Get("/about", h.About)
	mux.Get("/contact", h.Contact)
	mux.Get("/rooms", h.Rooms)
	mux.Get("/make-reservation", h.MakeReservation)
	mux.Get("/reservation-summary", h.ReservationSumm)
	mux.Get("/login", h.Login)

	mux.Post("/make-reservation", h.PostMakeReservation)
	mux.Post("/check-room-avail", h.PostCheckRoomAvail)
	mux.Post("/login", h.PostLogin)

	// add admin routes
	mux.Route("/admin", func(m chi.Router) {
		m.Use(mw.Authen)
		m.Get("/dashboard", h.AdminDashBoard)
	})

	// set up static file FileSystem relative to main.go
	fs := http.FileServer(http.Dir("../web/static"))

	// if not stripPrefix, when serve /static, go will look for /web/static/static
	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	return mux
}
