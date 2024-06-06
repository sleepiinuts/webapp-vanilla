package users

import "fmt"

type UserServ struct {
	ur UserRepos
}

func NewServ(ur UserRepos) *UserServ {
	return &UserServ{ur: ur}
}

func (us *UserServ) New(firstName, lastName, email, pwd, phone, role string) (int, error) {
	row, err := us.ur.new(firstName, lastName, email, pwd, phone, role)
	if err != nil {
		return -1, fmt.Errorf("[user] new: %w", err)
	}

	id := -1
	if err := row.Scan(&id); err != nil {
		return id, fmt.Errorf("[user] new - scan: %w", err)
	}

	return id, nil
}
