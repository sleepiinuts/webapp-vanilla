package rooms

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/qustavo/dotsql"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

var (
	ErrStmtNotFound = fmt.Errorf("sql statement not found")
	ErrStmtExec     = fmt.Errorf("unable to successfully query/exec stmt")
	ErrStructScan   = fmt.Errorf("unable to scan to struct")
)

type PostgresRoom struct {
	db  *sqlx.DB
	dot *dotsql.DotSql
	// logger *slog.Logger
}

// findById implements RoomRepos.
func (p *PostgresRoom) findById(id int) (*models.Room, error) {
	name := "findById"
	stmt, err := p.dot.Raw(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	var room models.Room
	err = p.db.QueryRowx(stmt, id).StructScan(&room)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	return &room, nil
}

func NewPostgresRoom(db *sqlx.DB, dot *dotsql.DotSql) *PostgresRoom {
	return &PostgresRoom{db: db, dot: dot}
}

// findAll implements RoomRepos.
func (p *PostgresRoom) findAll() ([]*models.Room, error) {
	name := "findAll"

	stmt, err := p.dot.Raw(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	var rooms []*models.Room

	rows, err := p.db.Queryx(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	defer rows.Close()
	for rows.Next() {
		var room models.Room
		err := rows.StructScan(&room)
		if err != nil {
			return nil, fmt.Errorf("%s: %w: %w", name, ErrStructScan, err)
		}

		rooms = append(rooms, &room)
	}

	return rooms, nil
}

var _ RoomRepos = &PostgresRoom{}
