package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/forms"
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
	ap *configs.AppProperties
}

func New(r *renders.Renderer, sm *scs.SessionManager, ap *configs.AppProperties) *Handler {
	return &Handler{r: r, sm: sm, ap: ap}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	h.r.RenderTemplateFromMap(w, r, "index.tmpl", &models.Template{})
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
	var (
		id  int
		err error
	)

	if r.URL.Query().Get("id") != "" {
		id, err = strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		h.sm.Put(r.Context(), "roomId", id)
		// h.ap.Logger.Info("StoreId", "roomId", h.sm.Get(r.Context(), "roomId"))
	} else {
		id = h.sm.PopInt(r.Context(), "roomId")
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
	// populate r.Form
	err := r.ParseForm()
	if err != nil {
		h.ap.Logger.Error("Invalid Form", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	form := forms.New(r.PostForm)
	form.Require("firstName", "lastName", "email")
	form.IsValidEmail("email")

	// parse session-message to template data
	if !form.IsValid() {
		w.WriteHeader(http.StatusForbidden)
		h.r.RenderTemplateFromMap(w, r, "make-reservation.tmpl",
			&models.Template{
				Form: r.PostForm, FormErrors: form.GetErrors(),
			})
		return
	}

	h.ap.Logger.Info("Make Reservation Success")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) PostCheckRoomAvail(w http.ResponseWriter, r *http.Request) {
	// parse session-message to template data
	// w.Write([]byte("datepicker: " + r.URL.Query().Get("datepicker")))

	// populate r.Form
	err := r.ParseForm()
	if err != nil {
		h.ap.Logger.Error("Invalid Form", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	form := forms.New(r.Form)

	// check for required fields
	form.Require("datepicker")
	if !form.IsValid() {

		id := fmt.Sprintf("%v", h.sm.Get(r.Context(), "roomId"))
		if id != "" {
			http.Redirect(w, r, fmt.Sprintf("/rooms?id=%s", id), http.StatusSeeOther)
		} else {
			h.ap.Logger.Error("Invalid RoomId", "id", id)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		return
	}

	h.ap.Logger.Info("Booking Success", "dateRange", r.FormValue("datepicker"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
