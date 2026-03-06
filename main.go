
package main

import (
	"context"
	"log"
	"os"
	"github.com/urfave/cli/v3"
	"funk/commands"
)



func main()  {
	cmd := &cli.Command{
		Name:  "funk",
		Usage: "suite of useful tools for pesky problems",
		Commands: []*cli.Command{
			commands.TimerCommand(),
		},
	}

	if err:=cmd.Run(context.Background(),os.Args); err !=nil {
		log.Fatal(err)
	}
}