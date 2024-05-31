package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/renders"
	"github.com/sleepiinuts/webapp-plain/test"
)

var (
	ap      *configs.AppProperties
	sm      *scs.SessionManager
	r       *renders.Renderer
	h       *Handler
	handler http.Handler
)

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

func init() {
	ap, sm = test.GetDependencies()
	r = renders.New(ap, sm, "../../web/templates/")
	h = New(r, sm, ap)

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

	handler = sm.LoadAndSave(mux)
}
