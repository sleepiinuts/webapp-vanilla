package configs

import (
	"html/template"
	"log/slog"
	"net/http"
)

type AppProperties struct {
	Tc       map[string]*template.Template
	UseCache bool
	Logger   *slog.Logger
	Cookies  *Cookies
}

type Cookies struct {
	HttpOnly bool
	Path     string
	SameSite http.SameSite
	Secure   bool
}

func New(tc map[string]*template.Template, usecache bool, logger *slog.Logger) *AppProperties {
	return &AppProperties{
		Tc: tc, UseCache: usecache, Logger: logger,
		Cookies: &Cookies{HttpOnly: true, Path: "/", SameSite: http.SameSiteLaxMode, Secure: false},
	}
}
