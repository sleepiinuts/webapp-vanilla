package models

type User struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Pwd       []byte `db:"password"`
	Role      string `db:"role"`
}
