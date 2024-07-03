package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
	"github.com/sleepiinuts/webapp-plain/pkg/repositories/reservations"
)

type calendar struct {
	roomID     int
	startDt    time.Time
	today      time.Time
	activities [][]*activity
}

type activity struct {
	date    time.Time
	resvID  int
	resvCSS string
}

var Header = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

func newCalendar(roomID int, sd, td time.Time, rs *reservations.ReservationServ) (*calendar, error) {
	if sd.IsZero() {

		tStr := td.Format(time.DateOnly)
		tStr = fmt.Sprintf("%s%s", tStr[:len(tStr)-2], "01")
		startDt, err := time.Parse(configs.DateFormat, tStr)
		if err != nil {
			panic("[new calendar] invalid date: " + td.String())
		}

		// calculate first calendar date
		sd = startDt.AddDate(0, 0, -int(startDt.Weekday()))
	}

	resvs, err := rs.FindByIdAndArrAndDep(roomID, sd, sd.Add(42*24*time.Hour))
	if err != nil {
		return nil, fmt.Errorf("newCalendar: %w", err)
	}

	c := &calendar{roomID: roomID, startDt: sd, today: td}
	c.generate(resvs)

	return c, nil
}

func (c *calendar) generate(resvs map[time.Time]*models.Reservation) {
	// return 6x7 calendar dates
	// NewCalendar guarantees startDt value exists

	var activities [][]*activity

	d := c.startDt
	for i := 0; i < 6; i++ {
		activities = append(activities, make([]*activity, 7))
		for j := 0; j < 7; j++ {
			activity := activity{}
			activity.date = d

			if resv, ok := resvs[d]; ok {
				resvCSS := ""

				switch {
				case d.Equal(resv.Arrival) && d.Equal(resv.Departure):
					// booking only 1 day
					resvCSS = "booked-only"
				case d.Equal(resv.Arrival) && d.Before(resv.Departure):
					// booking multiple days && is the first day of the trip
					resvCSS = "booked-head"
				case d.After(resv.Arrival) && d.Before(resv.Departure):
					// booking multiple days && is the intermediate day
					resvCSS = "booked-body"
				case d.Equal(resv.Departure):
					// booking multiple days && is the last day of the trip
					resvCSS = "booked-tail"
				default:
					// should not exist as the service generates only map of reservation within arr & dep dates
				}

				activity.resvID = resv.ID
				activity.resvCSS = resvCSS
			}

			activities[i][j] = &activity
			d = d.AddDate(0, 0, 1)
		}
	}

	// set activities to calendar
	c.activities = activities
}

func (a *activity) String() string {
	return fmt.Sprintf("{date:%s,resvID:%d,resvCSS:%s}",
		a.date.Format(configs.DateFormat), a.resvID, a.resvCSS)
}

func (a *activity) GetDate() string {
	return fmt.Sprintf("%02d", a.date.Day())
}

func (a *activity) GetResvCSS() string {
	return a.resvCSS
}

// func (c *calendar) String() string {
// 	str := fmt.Sprintf("{startDt:%s,today:%s}", c.startDt, c.today)
// 	calendar := c.Generate()
// 	for _, row := range calendar {
// 		str = fmt.Sprintf("%s\n%s", str, strings.Join(row, ","))
// 	}

// 	return str
// }

func (c *calendar) PrevStartDt() time.Time {
	return c.startDt.AddDate(0, -1, 0)
}

func (c *calendar) NextStartDt() time.Time {
	return c.startDt.AddDate(0, 1, 0)
}

func (c *calendar) GetStartDt() time.Time {
	return c.startDt
}

func (h *Handler) AdminDashBoard(w http.ResponseWriter, r *http.Request) {
	// get current year & month from "r"
	year := r.URL.Query().Get("y")
	month := r.URL.Query().Get("m")

	if year == "" || month == "" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		h.ap.Logger.Error("adminDashBoard", "year", year, "month", month)
		return
	}

	// get room id
	rid, err := strconv.Atoi(r.URL.Query().Get("rid"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		h.ap.Logger.Error("adminDashBoard: not specify roomID")
		return
	}

	sd, err := time.Parse(configs.DateFormat, fmt.Sprintf("%s-%s-01", year, month))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		h.ap.Logger.Error("adminDashBoard: startDt calculation", "error", err)
		return
	}
	// calculate first calendar date
	sd = sd.AddDate(0, 0, -int(sd.Weekday()))

	c, err := newCalendar(rid, sd, time.Now(), h.rs)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		h.ap.Logger.Error("adminDashBoard: newCalendar", "error", err)
		return
	}

	h.r.RenderTemplateFromMap(w, r, "admin-dashboard.tmpl", &models.Template{
		Data: map[string]any{"activities": c.activities},
	})
}
