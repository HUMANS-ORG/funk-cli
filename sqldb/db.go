package sqldb

import (
	"database/sql"
	"log"
	"fmt"
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
	defer db.Close()

	create := ` CREATE TABLE IF NOT EXISTS Timer (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		create_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        timer TEXT
    );`

	_,err := db.Exec(create)

	if err !=nil{
		log.Fatal(err)
	}

	log.Println("Table 'users' created successfully")

}

func Insert_data(h int,m int,s int)  {
	db := ConnectDB()

	defer db.Close()

	timer := fmt.Sprintf("%d:%d:%d", h, m, s)

	_,err :=db.Exec("INSERT INTO Timer(timer) VALUES(?)",timer)

	if err !=nil{
		log.Fatal(err)
	}

	log.Printf("insert successfully: %02d:%02d:%02d\n", h, m, s)
	
}

func Show_history()  {
	db := ConnectDB()
	defer db.Close()

	row,err := db.Query("SELECT id, create_at, timer FROM Timer ORDER BY id DESC")

	if err !=nil{
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next(){
		var id int
		var created string
		var timer string

		row.Scan(&id,&created,&timer)

		fmt.Printf("%d | %s | %s\n", id, created, timer)

	}
}