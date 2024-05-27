package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/internal/renders"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type Handler struct {
	r  *renders.Renderer
	sm *scs.SessionManager
}

func New(r *renders.Renderer, sm *scs.SessionManager) *Handler {
	return &Handler{r: r, sm: sm}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	h.r.RenderTemplateFromMap(w, "home.tmpl", &models.Template{})
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, "about.tmpl", &models.Template{})
}

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, "contact.tmpl", &models.Template{})
}

func (h *Handler) GradeSuperior(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, "grandsuperior.tmpl", &models.Template{})
}

func (h *Handler) DeluxeRoom(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, "deluxeroom.tmpl", &models.Template{})
}

func (h *Handler) MakeReservation(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, "make-reservation.tmpl", &models.Template{})
}
