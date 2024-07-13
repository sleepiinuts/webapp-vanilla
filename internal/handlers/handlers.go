package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/forms"
	"github.com/sleepiinuts/webapp-plain/internal/helpers"
	"github.com/sleepiinuts/webapp-plain/internal/renders"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
	"github.com/sleepiinuts/webapp-plain/pkg/repositories/reservations"
	"github.com/sleepiinuts/webapp-plain/pkg/repositories/restrictions"
	rms "github.com/sleepiinuts/webapp-plain/pkg/repositories/rooms"
	"github.com/sleepiinuts/webapp-plain/pkg/repositories/users"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/maps"
)

type Handler struct {
	r   *renders.Renderer
	sm  *scs.SessionManager
	ap  *configs.AppProperties
	rs  *reservations.ReservationServ
	rms *rms.RoomServ
	us  *users.UserServ
	rtn *restrictions.RestrictionServ
}

func New(r *renders.Renderer, sm *scs.SessionManager, ap *configs.AppProperties, rs *reservations.ReservationServ, rms *rms.RoomServ, us *users.UserServ, rtn *restrictions.RestrictionServ) *Handler {
	return &Handler{r: r, sm: sm, ap: ap, rs: rs, rms: rms, us: us, rtn: rtn}
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

	// enforce to get id
	if r.URL.Query().Get("id") != "" {
		id, err = strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			// TODO: more elegant way of handling error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		h.sm.Put(r.Context(), "roomId", id)
	} else {
		id = h.sm.PopInt(r.Context(), "roomId")
	}

	room, err := h.rms.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		h.ap.Logger.Error("[rooms] handler: findById", "error", err)
		return
	}

	// put room into session
	h.sm.Put(r.Context(), "rooms", map[int]models.Room{id: *room})

	// parse session-message to template data
	h.r.RenderTemplateFromMap(w, r, "rooms.tmpl", &models.Template{
		Data: map[string]any{
			"name": room.Name,
			"img":  room.ImgPath,
			"desc": room.Desc,
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

		// TODO: user redirect instead ??
		h.r.RenderTemplateFromMap(w, r, "make-reservation.tmpl",
			&models.Template{
				Form: r.PostForm, FormErrors: form.GetErrors(),
			})
		return
	}

	h.ap.Logger.Info("Make Reservation Success")
	h.sm.Put(r.Context(), "Flash", models.Flash{
		Body: "Make Reservation Success",
		Type: models.FTInfo,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) PostCheckRoomAvail(w http.ResponseWriter, r *http.Request) {
	// TODO: make the caller put room(s) into session

	// populate r.Form to get datepicker
	err := r.ParseForm()
	if err != nil {
		h.ap.Logger.Error("invalid Form", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	form := forms.New(r.Form)

	// check for required fields
	form.Require("datepicker")
	id := h.sm.Get(r.Context(), "roomId")

	redirect := func(id any) {

		if id == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		rid, err := strconv.Atoi(id.(string))

		if err != nil {
			h.ap.Logger.Error("invalid RoomId", "id", id)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, fmt.Sprintf("/rooms?id=%d", rid), http.StatusSeeOther)
		}
	}

	if !form.IsValid() {
		h.ap.Logger.Error("invalid form: missing date picker")

		redirect(id)
		return
	}

	dtRange := strings.Split(form.GetField("datepicker"), " - ")
	if len(dtRange) != 2 {
		h.ap.Logger.Error("invalid format datepicker", "datepicker", form.GetField("datepicker"))

		redirect(id)
		return
	}

	sd := helpers.MustParseTime(dtRange[0])
	ed := helpers.MustParseTime(dtRange[1])

	// TODO: get Room(s) from session
	// TODO: if !ok then put all rooms into session? --> meaning get redirect from homepage?
	rooms, ok := h.sm.Get(r.Context(), "rooms").(map[int]models.Room)
	if !ok {
		rooms, err = h.rms.FindAll()
		h.sm.Put(r.Context(), "rooms", rooms)

		if err != nil {
			h.ap.Logger.Error("[post-checkRoomAvail]handler", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// get room id(s)
	ids := maps.Keys(rooms)

	avlPeriod, err := h.rs.ListAvailRooms(ids, sd, ed)
	if err != nil {
		h.ap.Logger.Error("list avail rooms", "error", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	h.ap.Logger.Info("available period", "avlPeriod", avlPeriod)
	h.ap.Logger.Info("Booking Success", "dateRange", r.FormValue("datepicker"))

	// pass data to reservation summary page
	h.sm.Put(r.Context(), "datepicker", form.GetField("datepicker"))
	h.sm.Put(r.Context(), "avlPeriod", avlPeriod)

	// TODO: make reservation-summary page
	// redirect to xx page
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (h *Handler) ReservationSumm(w http.ResponseWriter, r *http.Request) {
	// get datepicker
	datePicker, ok := h.sm.Pop(r.Context(), "datepicker").(string)
	if !ok {
		h.ap.Logger.Error("datepicker not found")
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	// get avlPeriod
	avlPeriod, ok := h.sm.Pop(r.Context(), "avlPeriod").(map[int][]*reservations.Period)
	if !ok {
		h.ap.Logger.Error("avlPeriod not found")
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	// get Rooms
	rooms, ok := h.sm.Pop(r.Context(), "rooms").(map[int]models.Room)
	if !ok {
		h.ap.Logger.Error("[reservation-summ]handler", "error", "unable to get ROOMS from session")
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	h.ap.Logger.Debug("reservation-summ", "datepicker", datePicker, "avlPeriod", avlPeriod, "rooms", rooms)

	td := &models.Template{
		Data: map[string]any{
			"datepicker": datePicker,
			"avlPeriod":  avlPeriod,
			"rooms":      rooms,
		},
	}

	h.r.RenderTemplateFromMap(w, r, "reservation-summary.tmpl", td)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.r.RenderTemplateFromMap(w, r, "login.tmpl", &models.Template{})
}

func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
	// parse form
	err := r.ParseForm()
	if err != nil {
		h.ap.Logger.Error("[postLogin] invalid Form", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// check required fields
	form := forms.New(r.PostForm)
	form.Require("email", "password")
	if !form.IsValid() {
		w.WriteHeader(http.StatusForbidden)
		h.ap.Logger.Error("[postLogin] invalid Form", "err", "missing required field(s)")
		// TODO: make redirect
		return
	}

	// bcrypt pwd
	email := form.GetField("email")
	pwd := []byte(form.GetField("password"))

	// bpwd, err := bcrypt.GenerateFromPassword([]byte(pwd), configs.DefaultBcryptCost)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	h.ap.Logger.Error("[postLogin] hashing pwd", "error", err)
	// 	// TODO: make redirect
	// 	return
	// }

	// check if exist,user(email) && corresponding pwd
	id, dbPwd, err := h.us.Authen(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.ap.Logger.Error("[postLogin] authen", "error", err)
		// TODO: make redirect
		return
	}

	if err = bcrypt.CompareHashAndPassword(dbPwd, pwd); err != nil {
		h.ap.Logger.Error("[postLogin] invalid username/password", "error", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)

		// TODO: notify login fail
		return
	}
	// mitigate [OWASP] session fixation
	h.sm.RenewToken(r.Context())

	h.sm.Put(r.Context(), "userid", id)
	h.ap.Logger.Info("login success")

	// set default roomID & current month&year
	today := time.Now()
	route := fmt.Sprintf("/admin/dashboard?rid=1&y=%d&m=%02d", today.Year(), today.Month())

	http.Redirect(w, r, route, http.StatusSeeOther)
}
