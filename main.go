package main

import (
	"context"
	"log"
	"os"

	"funk/commands"

	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "funk",
		Usage: "Developer CLI tool",

		Commands: []*cli.Command{
			commands.FileEmptyCommand(),
			// commands.ConvertCommand()  <- teammates plug in here
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
