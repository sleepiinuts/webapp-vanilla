package users

import (
	"database/sql"
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

type PostgresUser struct {
	db  *sqlx.DB
	dot *dotsql.DotSql
}

func NewPostgresUser(db *sqlx.DB, dot *dotsql.DotSql) *PostgresUser {
	return &PostgresUser{db: db, dot: dot}
}

// authen implements UserRepos.
func (p *PostgresUser) authen(email string) (string, []byte, error) {
	name := "authen"
	id := ""
	pwd := []byte{}

	stmt, err := p.dot.Raw(name)
	if err != nil {
		return id, pwd, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	err = p.db.QueryRowx(stmt, email).Scan(&id, &pwd)
	if err == sql.ErrNoRows {
		// return pwd, nil
	} else if err != nil {
		return id, pwd, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	return id, pwd, nil
}

// findByEmail implements UserRepos.
func (p *PostgresUser) findByEmail(email string) (*models.User, error) {
	name := "findByEmail"
	stmt, err := p.dot.Raw(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	var u models.User
	err = p.db.QueryRowx(stmt, email).StructScan(&u)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	return &u, err
}

// new implements UserRepos.
func (p *PostgresUser) new(firstName, lastName, email, pwd, phone, role string) (int, error) {
	name := "new"
	stmt, err := p.dot.Raw(name)
	if err != nil {
		return -1, fmt.Errorf("%s: %w: %w", name, ErrStmtNotFound, err)
	}

	var id int
	err = p.db.QueryRowx(stmt, firstName, lastName, email, pwd, phone, role).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("%s: %w: %w", name, ErrStmtExec, err)
	}

	return id, err
}

var _ UserRepos = &PostgresUser{}
