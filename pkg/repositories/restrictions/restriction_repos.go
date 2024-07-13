package restrictions

import (
	"time"

	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type RestrictionRepos interface {
	findByIdAndEffAndExp(id int, eff, exp time.Time) (map[time.Time]*models.Restriction, error)
}
