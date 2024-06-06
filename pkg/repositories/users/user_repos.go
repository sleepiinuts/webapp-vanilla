package users

import "database/sql"

type UserRepos interface {
	new(firstName, lastName, email, pwd, phone, role string) (*sql.Row, error)
}
