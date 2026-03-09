package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "tasks.txt"

func AddTask(task string) {

	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	file.WriteString(task + "\n")

	fmt.Println("Task added")
}

func ListTasks() {

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("No tasks found")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 1
	for scanner.Scan() {
		fmt.Println(i, "-", scanner.Text())
		i++
	}
}

func DeleteTask(num string) {
	fmt.Println("Delete feature coming soon:", num)
}