package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func FileEmptyCommand() *cli.Command {
	return &cli.Command{
		Name:  "empty",
		Usage: "Find empty files and directories",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Value: ".",
				Usage: "Directory path to scan",
			},
		},
		Action: findEmpty,
	}
}

func findEmpty(ctx context.Context, c *cli.Command) error {
	path := c.String("path")
	FindEmptyFiles(path)
	return nil
}

func FindEmptyFiles(path string) {

	fmt.Println("Scanning for empty files and directories in:", path)

	foundAny := false

	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {

		if err != nil {
			fmt.Println("Warning: could not access:", p)
			return nil
		}

		// Skip the root path itself
		if p == path {
			return nil
		}

		if d.IsDir() {
			entries, readErr := os.ReadDir(p)
			if readErr != nil {
				fmt.Println("Warning: could not read dir:", p)
				return nil
			}
			if len(entries) == 0 {
				fmt.Println("Empty dir:", p)
				foundAny = true
			}
		} else {
			info, statErr := d.Info()
			if statErr != nil {
				fmt.Println("Warning: could not stat file:", p)
				return nil
			}
			if info.Size() == 0 {
				fmt.Println("Empty file:", p)
				foundAny = true
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error while scanning:", err)
		return
	}

	if !foundAny {
		fmt.Println("No empty files or directories found.")
	}
}
