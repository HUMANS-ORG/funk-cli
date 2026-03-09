 package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("add <task>")
		fmt.Println("list")
		fmt.Println("delete <number>")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		task := os.Args[2]
		AddTask(task)

	case "list":
		ListTasks()

	case "delete":
		DeleteTask(os.Args[2])

	default:
		fmt.Println("Unknown command")
	}
}