package users

import "fmt"

type UserServ struct {
	repos UserRepos
}

func NewServ(repos UserRepos) *UserServ {
	return &UserServ{repos: repos}
}

func (us *UserServ) New(firstName, lastName, email, pwd, phone, role string) (int, error) {
	row, err := us.repos.new(firstName, lastName, email, pwd, phone, role)
	if err != nil {
		return -1, fmt.Errorf("[user] new: %w", err)
	}

	id := -1
	if err := row.Scan(&id); err != nil {
		return id, fmt.Errorf("[user] new - scan: %w", err)
	}

	return id, nil
}
