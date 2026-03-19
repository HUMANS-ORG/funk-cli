 package main

import (
	"os"
	"github.com/urfave/cli/v3"
	"funk/todo"
	"log"
	"context"
)

func main() {
cmd := &cli.Command{
		Name:  "funk",
		Usage: "suite of useful tools for pesky problems",
		Commands: []*cli.Command{
		 todo.Todos(),
		},
	}

	if err:=cmd.Run(context.Background(),os.Args); err !=nil {
		log.Fatal(err)
	}
}