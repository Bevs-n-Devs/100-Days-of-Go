package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() error {
	var err error
	connStr := "postgres://username:password@localhost/akoto_tech?sslmode=disable" // this needs to updated to actual PostgreSQL DB credentials
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// make sure the connection is established
	return db.Ping()
}
