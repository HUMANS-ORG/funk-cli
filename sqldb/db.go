package sqldb

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db,err := sql.Open("sqlite3","./funk.db")
	if err != nil{
		log.Fatal(err)
	}

	return db
}

func Create_db() {
	db :=ConnectDB()
	create := ` CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );`

	_,err := db.Exec(create)

	if err !=nil{
		log.Fatal(err)
	}

	log.Println("Table 'users' created successfully")

}