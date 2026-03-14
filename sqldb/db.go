package sqldb

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/alexeyco/simpletable"
	
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
		create_at DATE DEFAULT (DATE('now')),
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
	fmt.Printf("\ninsert timer data in db successfully: %02d:%02d:%02d\n Command 'funk timer --his' ", h, m, s)
}

func Show_history()  {
	db := ConnectDB()
	defer db.Close()

	row,err := db.Query("SELECT id, DATE(create_at), timer FROM Timer ORDER BY id DESC")

	if err !=nil{
		log.Fatal(err)
	}

	defer row.Close()

	table :=simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter,Text: "ID"},
			{Align: simpletable.AlignCenter,Text: "CREATE DATE"},
			{Align: simpletable.AlignCenter,Text: "TIMER"},
		},	
	}

	for row.Next(){
		var id int
		var created string
		var timer string
		row.Scan(&id,&created,&timer)
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", id)},
			{Text: created},
			{Align: simpletable.AlignCenter, Text: timer},
		}
		table.Body.Cells = append(table.Body.Cells,r)
	}
	table.SetStyle(simpletable.StyleDefault)

	fmt.Println(table.String())
}

func  Delete_Record(index int)  {
	db := ConnectDB()
	defer db.Close()

	record,err:=db.Exec("DELETE FROM Timer WHERE id = ?",index)

	if err !=nil{
		log.Fatal(err)
	}
	row,err  := record.RowsAffected()

	if err !=nil{
		log.Fatal(err)
	}

	if row ==0 {
		fmt.Println("no record found or record already delete")
		return
	}

	fmt.Println("record delete successfully")
}