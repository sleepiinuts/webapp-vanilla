package renders

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

const basePath = "../../templates/"
const btPath = basePath + "base.tmpl"

var tc = make(map[string]*template.Template)
var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tPath := basePath + tmpl
	// btPath := basePath + "base.tmpl"

	parsedTmpl, err := template.ParseFiles(tPath, btPath)
	if err != nil {
		logger.Error("unable to find template: ", "path", tPath, "error", err)
		return
	}

	err = parsedTmpl.Execute(w, nil)
	if err != nil {
		logger.Error("unable to render template: ", "path", tPath, "error", err)
	}
}

func RenderTemplateFromMap(w http.ResponseWriter, tmpl string) {
	tPath := basePath + tmpl

	// cache template
	if _, ok := tc[tmpl]; !ok {

		// initial template
		page, err := template.New(tmpl).ParseFiles(basePath + tmpl)
		if err != nil {
			logger.Error("template initializing: ", "error", err)
			return
		}

		// check if layout exist
		pattern := "*.layout.tmpl"
		matches, err := filepath.Glob(basePath + pattern)
		if err != nil {
			logger.Error("template layout gathering: ", "basePath", basePath, "pattern", pattern)
			return
		}
		logger.Info("template layout matches: ", "matches", matches)

		// if layout exist, include layout into "initialized template (page)"
		var parsedTmpl *template.Template
		if len(matches) > 0 {
			parsedTmpl, err = page.ParseGlob(basePath + pattern)
			if err != nil {
				logger.Error("unable to find template: ", "path", tPath, "error", err)
				return
			}
		}

		logger.Info("caching template", "path", tPath)
		tc[tmpl] = parsedTmpl
	}

	err := tc[tmpl].Execute(w, nil)
	if err != nil {
		logger.Error("unable to render template: ", "templateName", tmpl, "error", err)
	}
}
