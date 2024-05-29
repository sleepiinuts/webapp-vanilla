package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/handlers"
	"github.com/sleepiinuts/webapp-plain/internal/renders"
	"github.com/sleepiinuts/webapp-plain/internal/routes"
)

const port = ":8080"

var (
	ap *configs.AppProperties
	r  *renders.Renderer
	h  *handlers.Handler
	sm *scs.SessionManager
)

func main() {
	// http.HandleFunc("/", h.Home)
	// http.HandleFunc("/about", h.About)

	ap.Logger.Info("Starting application", "port", port)
	http.ListenAndServe(port, sm.LoadAndSave(routes.Routes(h, ap)))
}

func init() {
	ap = configs.New(
		make(map[string]*template.Template),
		false,
		slog.New(slog.NewTextHandler(os.Stdout, nil)))

	// config session manager
	sm = scs.New()
	sm.Lifetime = 24 * time.Hour
	sm.Cookie.HttpOnly = ap.Cookies.HttpOnly
	sm.Cookie.Path = ap.Cookies.Path
	sm.Cookie.SameSite = ap.Cookies.SameSite
	sm.Cookie.Secure = ap.Cookies.Secure

	r = renders.New(ap)
	h = handlers.New(r, sm, ap)

}
