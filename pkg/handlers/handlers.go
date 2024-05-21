package handlers

import (
	"net/http"

	"github.com/sleepiinuts/webapp-plain/pkg/renders"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.tmpl")
}
