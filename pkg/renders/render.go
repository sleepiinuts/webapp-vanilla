package renders

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
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
		logger.Error("Unable to find template: ", "path", tPath, "error", err)
		return
	}

	err = parsedTmpl.Execute(w, nil)
	if err != nil {
		logger.Error("Unable to render template: ", "path", tPath, "error", err)
	}
}

func RenderTemplateFromMap(w http.ResponseWriter, tmpl string) {
	tPath := basePath + tmpl

	if _, ok := tc[tmpl]; !ok {
		parsedTmpl, err := template.ParseFiles(tPath, btPath)
		if err != nil {
			logger.Error("Unable to find template: ", "path", tPath, "error", err)
			return
		}
		logger.Info("Caching template", "path", tPath)
		tc[tmpl] = parsedTmpl
	}

	err := tc[tmpl].Execute(w, nil)
	if err != nil {
		logger.Error("Unable to render template: %s: %v", tPath, err)
	}
}
