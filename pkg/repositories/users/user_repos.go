package users

import "github.com/sleepiinuts/webapp-plain/pkg/models"

type UserRepos interface {
	new(firstName, lastName, email, pwd, phone, role string) (int, error)
	findByEmail(email string) (*models.User, error)
	authen(email string) (string, []byte, error)
}
