package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("sqlite3", "./funk.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}