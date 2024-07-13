package restrictions

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

type PostgresRestriction struct {
	db  *sqlx.DB
	dot *dotsql.DotSql
}

// findByIdAndEffAndExp implements RestrictionRepos.
func (p *PostgresRestriction) findByIdAndEffAndExp(id int, eff time.Time, exp time.Time) (map[time.Time]*models.Restriction, error) {
	name := "findByIdAndEffAndExp"
	stmt, err := p.dot.Raw(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	rests := make(map[time.Time]*models.Restriction)
	rows, err := p.db.Queryx(stmt, id, eff, exp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	defer rows.Close()
	for rows.Next() {
		var rest models.Restriction
		err := rows.StructScan(&rest)
		if err != nil {
			return nil, fmt.Errorf("%s: %w: %w", name, ErrStructScan, err)
		}

		st := rest.Effective
		en := rest.Expire

		for d := st; d.Before(en) || d.Equal(en); d = d.Add(24 * time.Hour) {
			rests[d] = &rest
		}

	}
	return rests, nil
}

func NewPostgresRestriction(db *sqlx.DB, dot *dotsql.DotSql) *PostgresRestriction {
	return &PostgresRestriction{db: db, dot: dot}
}

var _ RestrictionRepos = &PostgresRestriction{}
