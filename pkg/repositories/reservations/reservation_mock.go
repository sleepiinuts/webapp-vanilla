package reservations

import (
	"errors"
	"time"

	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type MockReservation struct{}

var ErrNotFound = errors.New("info not found")

// findByArrivalAndDeparture implements ReservationRepos.
func (m *MockReservation) findByArrivalAndDeparture(arr time.Time, dep time.Time) (map[int][]*models.Reservation, error) {
	a1 := mustParseTime("2024-06-01")
	d1 := mustParseTime("2024-06-12")

	if arr.Equal(a1) && dep.Equal(d1) {
		return map[int][]*models.Reservation{
			1: {
				{ID: 1, UserID: 1, RoomID: 1, Arrival: mustParseTime("2024-06-02"), Departure: mustParseTime("2024-06-04")},
				{ID: 3, UserID: 2, RoomID: 1, Arrival: mustParseTime("2024-06-06"), Departure: mustParseTime("2024-06-07")},
			},
			2: {
				{ID: 2, UserID: 1, RoomID: 2, Arrival: mustParseTime("2024-06-02"), Departure: mustParseTime("2024-06-04")},
			},
		}, nil
	}
	return map[int][]*models.Reservation{}, ErrNotFound
}

func mustParseTime(t string) time.Time {
	tt, err := time.Parse(configs.DateFormat, t)
	if err != nil {
		panic("unparsable time!")
	}

	return tt
}

var _ ReservationRepos = &MockReservation{}
