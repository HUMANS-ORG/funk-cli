package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v3"
)

func FileRecentCommand() *cli.Command {
	return &cli.Command{
		Name:  "recent",
		Usage: "Find recently modified files",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "days",
				Value: 1,
				Usage: "Show files modified within given days",
			},
			&cli.StringFlag{
				Name:  "path",
				Value: ".",
				Usage: "Directory path to scan",
			},
		},
		Action: findRecent,
	}
}

func findRecent(ctx context.Context, c *cli.Command) error {

	path := c.String("path")
	days := c.Int("days")

	fmt.Println("Scanning for files modified in last", days, "day(s)...")

	cutoff := time.Now().AddDate(0, 0, -days)
	foundAny := false

	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {

		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		if info.ModTime().After(cutoff) {
			fmt.Printf("Recent file: %s (modified: %s)\n", p, info.ModTime().Format("2006-01-02 15:04:05"))
			foundAny = true
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error while scanning:", err)
	}

	if !foundAny {
		fmt.Println("No recently modified files found.")
	}

	return nil
}
