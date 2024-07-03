package reservations

import (
	"time"

	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type ReservationRepos interface {
	// findByArrivalAndDeparture: find reservation(s) within given time period
	findByArrivalAndDeparture(arr, dep time.Time) (map[int][]*models.Reservation, error)
	findByIdAndArrAndDep(id int, arr, dep time.Time) (map[time.Time]*models.Reservation, error)
}
