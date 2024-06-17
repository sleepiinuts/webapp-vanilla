package test

import (
	"encoding/gob"
	"html/template"
	"log/slog"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
	"github.com/sleepiinuts/webapp-plain/pkg/repositories/reservations"
)

var (
	ap *configs.AppProperties
	sm *scs.SessionManager
)

type Cases struct {
	Name     string
	Req      Request
	Resp     Response
	IsExpErr bool
	ExpErr   error
}

type Request struct {
	Path        string
	Method      string
	QueryParams map[string]string
	PostForm    map[string]string
}
type Response struct {
	HttpStatus int
}

func GetDependencies() (*configs.AppProperties, *scs.SessionManager) {

	ap = configs.New(
		make(map[string]*template.Template),
		false,
		slog.New(slog.NewTextHandler(os.Stdout, nil)))

	// config session manager
	sm = scs.New()
	sm.Lifetime = 24 * time.Hour
	sm.Cookie.HttpOnly = ap.Cookies.HttpOnly
	sm.Cookie.Path = ap.Cookies.Path
	sm.Cookie.SameSite = ap.Cookies.SameSite
	sm.Cookie.Secure = ap.Cookies.Secure

	// register Flash model for encoding required in scs session
	gob.Register(models.Flash{})

	// register map[int][]*reservations.period
	gob.Register(map[int][]*reservations.Period{})

	// register Room model
	gob.Register(map[int]models.Room{})
	return ap, sm
}
