package main

import (
	"html/template"
	"log/slog"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	pathTmpl := "./templates/" + tmpl
	parsedTmpl, err := template.ParseFiles(pathTmpl)
	if err != nil {
		slog.Error("Unable to find template: %s: %v", pathTmpl, err)
		return
	}

	err = parsedTmpl.Execute(w, nil)
	if err != nil {
		slog.Error("Unable to render template: %s: %v", pathTmpl, err)
	}
}
