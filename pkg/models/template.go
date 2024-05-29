package models

import "net/url"

type Template struct {
	Data       map[string]any
	Form       url.Values
	FormErrors map[string][]string
	CSRFToken  string
	Flash      string
	Warning    string
	Error      string
}
