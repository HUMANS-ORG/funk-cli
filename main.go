package main

import (
	"context"
	"funk/commands"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"funk/commands"
)

func main() {
	cmd := &cli.Command{
		Name:  "funk",
		Usage: "suite of useful tools for pesky problems",
		Commands: []*cli.Command{
<<<<<<< HEAD
<<<<<<< HEAD
			commands.ConvertCommand(),
=======
>>>>>>> main
=======
>>>>>>> main
			commands.TimerCommand(),
		},
	}

	if err:=cmd.Run(context.Background(),os.Args); err !=nil {
		log.Fatal(err)
	}
}

