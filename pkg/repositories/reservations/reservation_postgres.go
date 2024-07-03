package reservations

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/qustavo/dotsql"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

var (
	ErrStmtNotFound = fmt.Errorf("sql statement not found")
	ErrStmtExec     = fmt.Errorf("unable to successfully query/exec stmt")
	ErrStructScan   = fmt.Errorf("unable to scan to struct")
)

type PostgresReservation struct {
	db  *sqlx.DB
	dot *dotsql.DotSql
}

// findByIdAndArrAndDep implements ReservationRepos.
func (p *PostgresReservation) findByIdAndArrAndDep(id int, arr time.Time, dep time.Time) (map[time.Time]*models.Reservation, error) {
	name := "findByIdAndArrAndDep"
	stmt, err := p.dot.Raw(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	resvs := make(map[time.Time]*models.Reservation)
	rows, err := p.db.Queryx(stmt, id, arr, dep)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	defer rows.Close()
	for rows.Next() {
		var resv models.Reservation
		err := rows.StructScan(&resv)
		if err != nil {
			return nil, fmt.Errorf("%s: %w: %w", name, ErrStructScan, err)
		}

		arr := resv.Arrival
		dep := resv.Departure

		for d := arr; d.Before(dep) || d.Equal(dep); d = d.Add(24 * time.Hour) {
			// it's guaranteed by the application that no overlapping booking is allowed
			resvs[d] = &resv
		}
	}

	return resvs, nil
}

// findByArrivalAndDeparture implements ReservationRepos.
func (p *PostgresReservation) findByArrivalAndDeparture(arr time.Time, dep time.Time) (map[int][]*models.Reservation, error) {
	name := "findByArrivalAndDeparture"
	stmt, err := p.dot.Raw(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	resvs := make(map[int][]*models.Reservation)
	rows, err := p.db.Queryx(stmt, arr, dep)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	defer rows.Close()
	for rows.Next() {
		var resv models.Reservation
		err := rows.StructScan(&resv)
		if err != nil {
			return nil, fmt.Errorf("%s: %w: %w", name, ErrStructScan, err)
		}

		resvs[resv.RoomID] = append(resvs[resv.RoomID], &resv)
	}

	return resvs, nil
}

func NewPostgresReservation(db *sqlx.DB, dot *dotsql.DotSql) *PostgresReservation {
	return &PostgresReservation{db: db, dot: dot}
}

var _ ReservationRepos = &PostgresReservation{}
