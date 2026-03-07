package main

import (
 "context"
 "log"
 "os"

 "github.com/urfave/cli/v3"
 "funk/commands"
)

func main() {

 cmd := &cli.Command{
  Name:  "funk",
  Usage: "Developer CLI tool",

  Commands: []*cli.Command{
   {
    Name:  "empty",
    Usage: "Find empty files",
    Flags: []cli.Flag{
     &cli.StringFlag{
      Name:  "path",
      Value: ".",
      Usage: "Directory path to scan",
     },
    },
    Action: func(ctx context.Context, c *cli.Command) error {

     path := c.String("path")
     commands.FindEmptyFiles(path)

     return nil
    },
   },
  },
 }

 if err := cmd.Run(context.Background(), os.Args); err != nil {
  log.Fatal(err)
 }
}