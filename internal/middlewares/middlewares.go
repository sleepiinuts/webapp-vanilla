package middlewares

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/sleepiinuts/webapp-plain/configs"
)

type Middlewares struct {
	ap *configs.AppProperties
}

func New(ap *configs.AppProperties) *Middlewares {
	return &Middlewares{ap: ap}
}

func (mw *Middlewares) CSRFHandler(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     mw.ap.Cookies.Path,
		HttpOnly: mw.ap.Cookies.HttpOnly,
		Secure:   mw.ap.Cookies.Secure,
		SameSite: mw.ap.Cookies.SameSite,
	})

	return csrfHandler
}
