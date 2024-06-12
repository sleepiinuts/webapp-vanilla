package configs

import (
	"html/template"
	"log/slog"
	"net/http"
	"time"
)

type DBConfig struct {
	Host string
	Port int
	User string
	Pwd  string
	DB   string
}

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

var DBConf = DBConfig{
	Host: "127.0.0.1",
	Port: 5432,
	User: "postgres",
	Pwd:  "P@ssw0rd",
	DB:   "go-webapp-vanilla",
}

var DateFormat = time.DateOnly

func New(tc map[string]*template.Template, usecache bool, logger *slog.Logger) *AppProperties {
	return &AppProperties{
		Tc: tc, UseCache: usecache, Logger: logger,
		Cookies: &Cookies{HttpOnly: true, Path: "/", SameSite: http.SameSiteLaxMode, Secure: false},
	}
}
