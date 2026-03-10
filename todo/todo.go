 package todo

import "funk/database"

func InitTable() {

	db := database.ConnectDB()

	db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT
	)
	`)
}
func AddTask(task string) {

	db := database.ConnectDB()

	stmt, _ := db.Prepare("INSERT INTO tasks(task) VALUES(?)")

	stmt.Exec(task)
}
func ListTasks() {

	db := database.ConnectDB()

	rows, _ := db.Query("SELECT id, task FROM tasks")

	for rows.Next() {

		var id int
		var task string

		rows.Scan(&id, &task)

		println(id, "-", task)
	}
}