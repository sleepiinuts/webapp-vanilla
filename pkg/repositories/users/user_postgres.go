package users

import (
	"database/sql"
	"log/slog"

	"github.com/qustavo/dotsql"
)

type PostgresUser struct {
	db     *sql.DB
	dot    *dotsql.DotSql
	logger *slog.Logger
}

func New(db *sql.DB, dot *dotsql.DotSql, logger *slog.Logger) *PostgresUser {
	return &PostgresUser{db: db, dot: dot, logger: logger}
}

// new implements UserRepos.
func (p *PostgresUser) new(firstName, lastName, email, pwd, phone, role string) (*sql.Row, error) {
	return p.dot.QueryRow(p.db, "new", firstName, lastName, email, pwd, phone, role)
}

var _ UserRepos = &PostgresUser{}
