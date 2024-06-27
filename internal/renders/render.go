package renders

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

// const basePath = "../web/templates/"

// const btPath = basePath + "base.tmpl"

// var tc = make(map[string]*template.Template)
// var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

var funcMap = template.FuncMap{
	"isEmptyFlash": isEmptyFlash,
	"isLoggedIn":   isLoggedIn,
}

func isEmptyFlash(f models.Flash) bool {
	return f == models.Flash{}
}

func isLoggedIn(data map[string]any) bool {
	if data == nil {
		return false
	}

	if _, ok := data["isLoggedIn"]; !ok {
		return false
	}
	return true
}

type Renderer struct {
	ap       *configs.AppProperties
	sm       *scs.SessionManager
	basePath string
}

var (
	ErrInitTmpl        = errors.New("inti template")
	ErrFindTmplLayout  = errors.New("find template layout")
	ErrParseTmplLayout = errors.New("parse template layout")
	ErrParseFlash      = errors.New("parse flash data")
	ErrExeTmpl         = errors.New("execute template")
)

func New(ap *configs.AppProperties, sm *scs.SessionManager, basePath string) *Renderer {
	return &Renderer{ap: ap, sm: sm, basePath: basePath}
}

func (r *Renderer) RenderTemplateFromMap(w http.ResponseWriter, rq *http.Request, tmpl string, td *models.Template) error {
	tPath := r.basePath + tmpl

	// set default template data
	td.CSRFToken = nosurf.Token(rq)
	// r.ap.Logger.Info("CSRF Token", "token", td.CSRFToken)

	// cache template
	if _, ok := r.ap.Tc[tmpl]; !ok || !r.ap.UseCache {

		// initial template
		page, err := template.New(tmpl).Funcs(funcMap).ParseFiles(r.basePath + tmpl)
		if err != nil {
			r.ap.Logger.Error("template init: ", "error", err)
			return fmt.Errorf("%w:%w", ErrInitTmpl, err)
		}

		// check if layout exist
		pattern := "*.layout.tmpl"
		matches, err := filepath.Glob(r.basePath + pattern)
		if err != nil {
			r.ap.Logger.Error("template layout gathering: ", "basePath", r.basePath, "pattern", pattern)
			return fmt.Errorf("%w:%w", ErrFindTmplLayout, err)
		}
		r.ap.Logger.Info("template layout matches: ", "matches", matches)

		// if layout exist, include layout into "initialized template (page)"
		var parsedTmpl *template.Template
		if len(matches) > 0 {
			parsedTmpl, err = page.ParseGlob(r.basePath + pattern)
			if err != nil {
				r.ap.Logger.Error("unable to find template: ", "path", tPath, "error", err)
				return fmt.Errorf("%w:%w", ErrParseTmplLayout, err)
			}
		}

		r.ap.Logger.Info("caching template", "path", tPath)
		r.ap.Tc[tmpl] = parsedTmpl
	}

	// get flash info if any
	if r.sm.Get(rq.Context(), "Flash") != nil {
		flash, ok := r.sm.Pop(rq.Context(), "Flash").(models.Flash)
		if !ok {
			r.ap.Logger.Error("Flash casting")
			return ErrParseFlash
		} else {
			td.Flash = flash
		}
	}

	// TODO: get userid from session, if exist means already logged in
	if r.sm.Get(rq.Context(), "userid") != nil {
		if td.Data == nil {
			td.Data = make(map[string]any)
		}

		td.Data["isLoggedIn"] = true
	}

	err := r.ap.Tc[tmpl].Execute(w, td)
	if err != nil {
		r.ap.Logger.Error("unable to render template: ", "templateName", tmpl, "error", err)
		return fmt.Errorf("%w:%w", ErrExeTmpl, err)
	}

	return nil
}
