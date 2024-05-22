package models

type Template struct {
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
