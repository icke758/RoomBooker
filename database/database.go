package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "rooms.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
