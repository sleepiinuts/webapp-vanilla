package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/pkg/handlers"
	"github.com/sleepiinuts/webapp-plain/pkg/renders"
)

const port = ":8080"

var (
	ap *configs.AppProperties
	r  *renders.Renderer
	h  *handlers.Handler
)

func main() {
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/about", h.About)

	ap.Logger.Info("Starting application", "port", port)
	http.ListenAndServe(port, nil)
}

func init() {
	ap = configs.New(
		make(map[string]*template.Template),
		true,
		slog.New(slog.NewTextHandler(os.Stdout, nil)))

	r = renders.New(ap)
	h = handlers.New(r)
}
