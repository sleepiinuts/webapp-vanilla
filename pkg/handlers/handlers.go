package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
	"github.com/sleepiinuts/webapp-plain/pkg/renders"
)

type Handler struct {
	r  *renders.Renderer
	sm *scs.SessionManager
}

func New(r *renders.Renderer, sm *scs.SessionManager) *Handler {
	return &Handler{r: r, sm: sm}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// working with scs session
	h.sm.Put(r.Context(), "ip", r.RemoteAddr)

	h.r.RenderTemplateFromMap(w, "home.tmpl", &models.Template{})
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	// get session-message
	ip := h.sm.GetString(r.Context(), "ip")

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, "about.tmpl", &models.Template{
		Data:  map[string]any{"ip": ip},
		Flash: "Greeting!!",
	})
}
