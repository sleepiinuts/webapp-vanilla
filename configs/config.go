package configs

import (
	"html/template"
	"log/slog"
)

type AppProperties struct {
	Tc       map[string]*template.Template
	UseCache bool
	Logger   *slog.Logger
}

func New(tc map[string]*template.Template, usecache bool, logger *slog.Logger) *AppProperties {
	return &AppProperties{Tc: tc, UseCache: usecache, Logger: logger}
}
