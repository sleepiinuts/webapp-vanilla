package restrictions

import (
	"time"

	"github.com/sleepiinuts/webapp-plain/pkg/models"
)

type RestrictionServ struct {
	repos RestrictionRepos
}

func New(repos RestrictionRepos) *RestrictionServ {
	return &RestrictionServ{repos: repos}
}

func (rs *RestrictionServ) FindByIdAndEffAndExp(id int, eff, exp time.Time) (map[time.Time]*models.Restriction, error) {
	return rs.repos.findByIdAndEffAndExp(id, eff, exp)
}
