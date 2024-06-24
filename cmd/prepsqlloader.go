package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/qustavo/dotsql"
)

func prepSqlLoader() {
	dots = make(map[string]*dotsql.DotSql)
	basePath, _ := os.Getwd()
	basePath = filepath.Dir(basePath) + "/pkg/repositories"

	if dot, err := dotsql.LoadFromFile(basePath + "/rooms/room.sql"); true {
		if err != nil {
			log.Fatal("room sql loader error: ", err)
		}

		dots["room"] = dot
	}

	if dot, err := dotsql.LoadFromFile(basePath + "/reservations/reservation.sql"); true {
		if err != nil {
			log.Fatal("reservation sql loader error: ", err)
		}

		dots["reservation"] = dot
	}

	if dot, err := dotsql.LoadFromFile(basePath + "/users/user.sql"); true {
		if err != nil {
			log.Fatal("user sql loader error: ", err)
		}

		dots["user"] = dot
	}
}
