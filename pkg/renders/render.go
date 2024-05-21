package renders

import (
	"html/template"
	"log/slog"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	pathTmpl := "../../templates/" + tmpl
	parsedTmpl, err := template.ParseFiles(pathTmpl)
	if err != nil {
		slog.Error("Unable to find template: %s: %v", pathTmpl, err)
		return
	}

	err = parsedTmpl.Execute(w, nil)
	if err != nil {
		slog.Error("Unable to render template: %s: %v", pathTmpl, err)
	}
}
