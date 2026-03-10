 package main

import (
	"fmt"
	"os"

	"funk/todo"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: funk todo")
		return
	}

	switch os.Args[1] {

	case "todo":
		todo.InitTable()

	default:
		fmt.Println("Command not found")
	}
}