package renders

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

const basePath = "../web/templates/"
const btPath = basePath + "base.tmpl"

// var tc = make(map[string]*template.Template)
// var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

var funcMap = template.FuncMap{
	"isEmptyFlash": isEmptyFlash,
}

func isEmptyFlash(f models.Flash) bool {
	return f == models.Flash{}
}

type Renderer struct {
	ap *configs.AppProperties
	sm *scs.SessionManager
}

func New(ap *configs.AppProperties, sm *scs.SessionManager) *Renderer {
	return &Renderer{ap: ap, sm: sm}
}

func (r *Renderer) RenderTemplate(w http.ResponseWriter, tmpl string) {
	tPath := basePath + tmpl
	// btPath := basePath + "base.tmpl"

	parsedTmpl, err := template.ParseFiles(tPath, btPath)
	if err != nil {
		r.ap.Logger.Error("unable to find template: ", "path", tPath, "error", err)
		return
	}

	err = parsedTmpl.Execute(w, nil)
	if err != nil {
		r.ap.Logger.Error("unable to render template: ", "path", tPath, "error", err)
	}
}

func (r *Renderer) RenderTemplateFromMap(w http.ResponseWriter, rq *http.Request, tmpl string, td *models.Template) {
	tPath := basePath + tmpl

	// set default template data
	td.CSRFToken = nosurf.Token(rq)
	// r.ap.Logger.Info("CSRF Token", "token", td.CSRFToken)

	// cache template
	if _, ok := r.ap.Tc[tmpl]; !ok || !r.ap.UseCache {

		// initial template
		page, err := template.New(tmpl).Funcs(funcMap).ParseFiles(basePath + tmpl)
		if err != nil {
			r.ap.Logger.Error("template initializing: ", "error", err)
			return
		}

		// check if layout exist
		pattern := "*.layout.tmpl"
		matches, err := filepath.Glob(basePath + pattern)
		if err != nil {
			r.ap.Logger.Error("template layout gathering: ", "basePath", basePath, "pattern", pattern)
			return
		}
		r.ap.Logger.Info("template layout matches: ", "matches", matches)

		// if layout exist, include layout into "initialized template (page)"
		var parsedTmpl *template.Template
		if len(matches) > 0 {
			parsedTmpl, err = page.ParseGlob(basePath + pattern)
			if err != nil {
				r.ap.Logger.Error("unable to find template: ", "path", tPath, "error", err)
				return
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
		} else {
			td.Flash = flash
		}
	}

	err := r.ap.Tc[tmpl].Execute(w, td)
	if err != nil {
		r.ap.Logger.Error("unable to render template: ", "templateName", tmpl, "error", err)
	}
}
