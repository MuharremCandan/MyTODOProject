package config

import (
	"database/sql"
)

var connString string = "postgre://postgres:password@localhost/postgres?sslmode=disable"

func GetDB() (db *sql.DB, err error) {

	db, err = sql.Open("postgres", connString)

	if err != nil {
		return
	}
	statemet, _ := db.Prepare("CREATE TABLE IF NOT EXISTS Todos (Id INTEGER PRIMARY KEY, Detail TEXT, Completed BIT);")
	statemet.Exec()
	return

}
