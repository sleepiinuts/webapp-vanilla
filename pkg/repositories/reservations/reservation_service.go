package reservations

import (
	"fmt"
	"time"

	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type ReservationServ struct {
	repo ReservationRepos
}

type Period struct {
	From time.Time
	To   time.Time
}

func New(repo ReservationRepos) *ReservationServ {
	return &ReservationServ{repo: repo}
}

func (rs *ReservationServ) ListAvailRooms(arr, dep time.Time) (map[int][]*Period, error) {
	// TODO: make inq all room service
	allRooms := []int{1, 2, 3, 4}

	bookedRooms, err := rs.repo.findByArrivalAndDeparture(arr, dep)
	availRooms := make(map[int][]*Period)

	inquiredPeriod := &Period{From: arr, To: dep}
	if err != nil {
		return map[int][]*Period{}, fmt.Errorf("listAvailRooms: %w", err)
	}

	for _, id := range allRooms {

		if reserv, ok := bookedRooms[id]; !ok {
			availRooms[id] = []*Period{inquiredPeriod}
		} else {
			availRooms[id] = disectPeriod(*inquiredPeriod, reserv)
		}
	}

	return availRooms, nil
}

func disectPeriod(inqP Period, rs []*models.Reservation) []*Period {
	ps := []*Period{}

	var lowerBound time.Time

	idx := -1
	for _, r := range rs {
		if inqP.To.Sub(inqP.From) <= 0 {
			break
		}
		idx++
		lowerBound = inqP.From

		if r.Arrival.Sub(lowerBound) > 0 {
			lowerBound = r.Arrival
		}

		if lowerBound.Sub(inqP.From) > 0 {
			ps = append(ps, &Period{From: inqP.From, To: lowerBound.AddDate(0, 0, -1)})
		}

		inqP.From = r.Departure.AddDate(0, 0, 1)
	}

	if idx != -1 && inqP.To.Sub(rs[idx].Departure) > 0 {
		ps = append(ps, &Period{From: rs[idx].Departure.AddDate(0, 0, 1), To: inqP.To})
	}

	return ps
}
