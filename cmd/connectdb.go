package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/internal/helpers"
)

func connectDB() *sql.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		configs.DBConf.User, configs.DBConf.Pwd,
		configs.DBConf.Host, configs.DBConf.Port,
		configs.DBConf.DB)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		helpers.Fatal("connection refuse", err, ap.Logger)
	}

	err = db.Ping()
	if err != nil {
		helpers.Fatal("ping refuse", err, ap.Logger)
	}

	ap.Logger.Info("connection establish!")
	return db
}
