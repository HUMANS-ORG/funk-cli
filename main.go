package main

import (
	"fmt"
	"os"

	"funk/todo"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: funk <tool>")
		return
	}

	switch os.Args[1] {

	case "todo":
		todo.Run(os.Args[2:])

	default:
		fmt.Println("Tool not found")
	}
}