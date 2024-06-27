package users

import "github.com/sleepiinuts/webapp-plain/pkg/models"

type UserServ struct {
	repos UserRepos
}

func New(repos UserRepos) *UserServ {
	return &UserServ{repos: repos}
}

func (us *UserServ) New(firstName, lastName, email, pwd, phone, role string) (int, error) {
	return us.repos.new(firstName, lastName, email, pwd, phone, role)
}

func (us *UserServ) Authen(email string) (string, []byte, error) {
	// return id,password,role[tbc]
	return us.repos.authen(email)
}

func (us *UserServ) FindByEmail(email string) (*models.User, error) {
	return us.repos.findByEmail(email)
}
