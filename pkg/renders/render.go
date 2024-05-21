package renders

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/sleepiinuts/webapp-plain/configs"
)

const basePath = "../../templates/"
const btPath = basePath + "base.tmpl"

// var tc = make(map[string]*template.Template)
// var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type Renderer struct {
	ap *configs.AppProperties
}

func New(ap *configs.AppProperties) *Renderer {
	return &Renderer{ap: ap}
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

func (r *Renderer) RenderTemplateFromMap(w http.ResponseWriter, tmpl string) {
	tPath := basePath + tmpl

	// cache template
	if _, ok := r.ap.Tc[tmpl]; !ok {

		// initial template
		page, err := template.New(tmpl).ParseFiles(basePath + tmpl)
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

	err := r.ap.Tc[tmpl].Execute(w, nil)
	if err != nil {
		r.ap.Logger.Error("unable to render template: ", "templateName", tmpl, "error", err)
	}
}
