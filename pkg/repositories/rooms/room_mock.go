package rooms

import (
	"fmt"

	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type MockRoom struct{}

var rooms = map[int]models.Room{
	1: {ID: 1, Name: "Grand Superiod", Desc: "sample grand superior description",
		Price: 100.5, ImgPath: "../static/images/grandsuperior.png"},
	2: {ID: 2, Name: "Deluxe Room", Desc: "sample deluxe room description",
		Price: 80.75, ImgPath: "../static/images/deluxeroom.png"},
}

// findAll implements RoomRepos.
func (m *MockRoom) findAll() (map[int]models.Room, error) {
	return rooms, nil
}

// findById implements RoomRepos.
func (m *MockRoom) findById(id int) (*models.Room, error) {
	room, ok := rooms[id]

	if !ok {
		return nil, fmt.Errorf("room not found")
	}

	return &room, nil
}

var _ RoomRepos = &MockRoom{}
