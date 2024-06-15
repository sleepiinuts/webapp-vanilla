package rooms

import "github.com/sleepiinuts/webapp-plain/pkg/models"

type RoomRepos interface {
	findAll() ([]*models.Room, error)
}
