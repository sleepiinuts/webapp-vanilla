package middlewares

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"github.com/sleepiinuts/webapp-plain/configs"
)

type Middlewares struct {
	ap *configs.AppProperties
	sm *scs.SessionManager
}

func New(ap *configs.AppProperties, sm *scs.SessionManager) *Middlewares {
	return &Middlewares{ap: ap, sm: sm}
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

// TODO: remove userServ
func (mw *Middlewares) Authen(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := mw.sm.Get(r.Context(), "userid")

		if id == nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}
