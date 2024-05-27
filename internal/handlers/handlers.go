package handlers

import (
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/internal/renders"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type RoomsInfo struct {
	name string
	img  string
}

var rooms = map[int]RoomsInfo{
	1: {name: "Grand Superior", img: "../static/images/grandsuperior.png"},
	2: {name: "Deluxe Room", img: "../static/images/deluxeroom.png"},
}

type Handler struct {
	r  *renders.Renderer
	sm *scs.SessionManager
}

func New(r *renders.Renderer, sm *scs.SessionManager) *Handler {
	return &Handler{r: r, sm: sm}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	h.r.RenderTemplateFromMap(w, r, "home.tmpl", &models.Template{})
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, r, "about.tmpl", &models.Template{})
}

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, r, "contact.tmpl", &models.Template{})
}

func (h *Handler) Rooms(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, ok := rooms[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, r, "rooms.tmpl", &models.Template{
		Data: map[string]any{
			"name": rooms[id].name,
			"img":  rooms[id].img,
		}})
}

func (h *Handler) MakeReservation(w http.ResponseWriter, r *http.Request) {

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, r, "make-reservation.tmpl", &models.Template{})
}

func (h *Handler) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, r, "make-reservation.tmpl", &models.Template{})
}
