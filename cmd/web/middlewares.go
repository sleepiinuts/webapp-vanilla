package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func csrfHandler(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     ap.Cookies.Path,
		HttpOnly: ap.Cookies.HttpOnly,
		Secure:   ap.Cookies.Secure,
		SameSite: ap.Cookies.SameSite,
	})

	return csrfHandler
}
