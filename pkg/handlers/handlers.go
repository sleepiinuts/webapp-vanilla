package handlers

import (
	"net/http"

	"github.com/sleepiinuts/webapp-plain/pkg/renders"
)

type Handler struct {
	r *renders.Renderer
}

func New(r *renders.Renderer) *Handler {
	return &Handler{r: r}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.r.RenderTemplateFromMap(w, "home.tmpl")
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	h.r.RenderTemplateFromMap(w, "about.tmpl")
}
