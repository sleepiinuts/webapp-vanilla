package rooms

import "github.com/sleepiinuts/webapp-plain/pkg/models"

type RoomServ struct {
	repos RoomRepos
}

func New(repos RoomRepos) *RoomServ {
	return &RoomServ{repos: repos}
}

func (rs *RoomServ) FindAll() ([]*models.Room, error) {
	return rs.repos.findAll()
}

func (rs *RoomServ) FindById(id int) (*models.Room, error) {
	return rs.repos.findById(id)
}
