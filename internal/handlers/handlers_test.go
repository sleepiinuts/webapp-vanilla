package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/renders"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
	"github.com/sleepiinuts/webapp-plain/pkg/repositories/reservations"
	"github.com/sleepiinuts/webapp-plain/pkg/repositories/rooms"
	"github.com/sleepiinuts/webapp-plain/test"
)

var (
	ap      *configs.AppProperties
	sm      *scs.SessionManager
	r       *renders.Renderer
	h       *Handler
	handler http.Handler
)

type sessionValues map[string]any

// const (
// 	contentTypeJSON = "application/json"
// 	contentTypeForm = "application/x-www-form-urlencoded"
// )

func TestHandlers(t *testing.T) {
	cases := []test.Cases{
		{
			Name: "[Home]Valid",
			Req: test.Request{
				Path:   "/",
				Method: "GET",
			},
			Resp: test.Response{
				HttpStatus: http.StatusOK,
			},
		},
		{
			Name: "[Home]Invalid-method",
			Req: test.Request{
				Path:   "/",
				Method: "POST",
			},
			Resp: test.Response{
				HttpStatus: http.StatusMethodNotAllowed,
			},
		},
		{
			Name: "[make-reservation]valid with body-form",
			Req: test.Request{
				Path:   "/make-reservation",
				Method: "POST",
				PostForm: map[string]string{
					"firstName":   "nut",
					"lastName":    "paw",
					"email":       "abc@gmail.com",
					"phoneNumber": "xx-xxx-xxxx",
				},
			},
			Resp: test.Response{
				HttpStatus: http.StatusOK,
			},
		},
		{
			Name: "[rooms]invalid no roomid",
			Req: test.Request{
				Path:   "/rooms",
				Method: "GET",
			},
			Resp: test.Response{
				HttpStatus: http.StatusNotFound,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			// rq := httptest.NewRequest(c.Req.Method, c.Req.Path, nil)
			// w := httptest.NewRecorder()

			ts := httptest.NewServer(handler)
			defer ts.Close()

			var (
				resp *http.Response
				err  error
			)

			if c.Req.Method == "GET" {
				resp, err = ts.Client().Get(ts.URL + c.Req.Path)
			} else {
				// resp, err = ts.Client().Post(ts.URL+c.Req.Path, contentTypeForm, nil)
				values := url.Values{}
				for k, v := range c.Req.PostForm {
					values.Add(k, v)
				}

				resp, err = ts.Client().PostForm(ts.URL+c.Req.Path, values)
			}

			if err != nil {
				t.Fail()
				t.Logf("unexpected error: got %v\n", err)
			}

			defer resp.Body.Close()

			// h.Home(w, rq)
			if resp.StatusCode != c.Resp.HttpStatus {
				t.Fail()
				t.Logf("mismatch statuse code, expected: %d, but got %d\n",
					c.Resp.HttpStatus, resp.StatusCode)
			}
		})
	}
}

func TestPostCheckRoomAvail(t *testing.T) {
	// TODO: add redirect location??
	var cases = []struct {
		name           string
		genUrlValues   func() url.Values
		smv            map[string]any
		expectedStatus int
		dstLoc         string
	}{
		{
			name: "happy case",
			genUrlValues: func() url.Values {
				values := url.Values{}
				values.Add("datepicker", "2024-06-01 - 2024-06-12")
				return values
			},
			smv: map[string]any{
				"roomId": "1",
				"rooms": map[int]models.Room{
					1: {ID: 1, Name: "Grand Superiod", Desc: "sample grand superior description",
						Price: 100.5, ImgPath: "../static/images/grandsuperior.png"},
					2: {ID: 2, Name: "Deluxe Room", Desc: "sample deluxe room description",
						Price: 80.75, ImgPath: "../static/images/deluxeroom.png"},
				},
			},
			dstLoc:         "/reservation-summary",
			expectedStatus: http.StatusSeeOther,
		},
		{
			name:           "no form parsed",
			genUrlValues:   func() url.Values { return nil },
			smv:            nil,
			expectedStatus: http.StatusSeeOther,
			dstLoc:         "/",
		},
		{
			name:         "no form parsed & roomId not a number",
			genUrlValues: func() url.Values { return nil },
			smv: map[string]any{
				"roomId": "abc",
			},
			expectedStatus: http.StatusSeeOther,
			dstLoc:         "/",
		},
		{
			name:         "no form parsed & with roomId",
			genUrlValues: func() url.Values { return nil },
			smv: map[string]any{
				"roomId": "1",
			},
			expectedStatus: http.StatusSeeOther,
			dstLoc:         "/rooms?id=1",
		},
		{
			name: "datepicker split error",
			genUrlValues: func() url.Values {
				values := url.Values{}
				values.Add("datepicker", "2024-06-01")
				return values
			},
			smv: map[string]any{
				"roomId": "1",
				"rooms": map[int]models.Room{
					1: {ID: 1, Name: "Grand Superiod", Desc: "sample grand superior description",
						Price: 100.5, ImgPath: "../static/images/grandsuperior.png"},
					2: {ID: 2, Name: "Deluxe Room", Desc: "sample deluxe room description",
						Price: 80.75, ImgPath: "../static/images/deluxeroom.png"},
				},
			},
			dstLoc:         "/rooms?id=1",
			expectedStatus: http.StatusSeeOther,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// set up form: datepicker
			req := httptest.NewRequest(http.MethodPost, "/check-roomm-avail", strings.NewReader(c.genUrlValues().Encode()))
			req.Header.Set("Content-type", "application/x-www-form-urlencoded")

			rec := httptest.NewRecorder()
			mockSessionValues(sm, c.smv)(http.HandlerFunc(h.PostCheckRoomAvail)).ServeHTTP(rec, req)

			got := rec.Result()
			if got.StatusCode != c.expectedStatus {
				t.Fail()
				t.Logf("unexpected status code: expected %d, but got %d\n", c.expectedStatus, got.StatusCode)
			}

			if c.dstLoc != "" {
				if c.dstLoc != got.Header.Get("Location") {
					t.Fail()
					t.Logf("expected loc: %s, but got %s\n", c.dstLoc, got.Header.Get("Location"))
				}
			}

		})
	}
}

func mockSessionValues(sm *scs.SessionManager, smv sessionValues) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return sm.LoadAndSave(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for k, v := range smv {
				sm.Put(r.Context(), k, v)
			}

			next.ServeHTTP(w, r)

			// // clean up session
			// for k := range smv {
			// 	sm.Pop(r.Context(), k)
			// }
			sm.Clear(r.Context())
		}))
	}
}
func init() {
	ap, sm = test.GetDependencies()
	r = renders.New(ap, sm, "../../web/templates/")

	// services
	rs := reservations.New(&reservations.MockReservation{})
	rms := rooms.New(&rooms.MockRoom{})

	h = New(r, sm, ap, rs, rms, nil)

	// routes.Routes(h,ap)
	mux := chi.NewRouter()

	// create middleware -- omit for testing handler purpose
	// mw := middlewares.New(ap)

	// mux.Use(mw.CSRFHandler)

	mux.Get("/", h.Home)
	mux.Get("/about", h.About)
	mux.Get("/contact", h.Contact)
	mux.Get("/rooms", h.Rooms)
	mux.Get("/make-reservation", h.MakeReservation)

	mux.Post("/make-reservation", h.PostMakeReservation)
	mux.Post("/check-room-avail", h.PostCheckRoomAvail)

	// set up static file FileSystem relative to main.go
	fs := http.FileServer(http.Dir("../web/static"))

	// if not stripPrefix, when serve /static, go will look for /web/static/static
	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	// setup handler
	handler = sm.LoadAndSave(mux)
}
