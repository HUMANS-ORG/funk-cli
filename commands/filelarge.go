package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func FileLargeCommand() *cli.Command {
	return &cli.Command{
		Name:  "large",
		Usage: "Find files larger than a given size in GB",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "size",
				Value: 1,
				Usage: "Minimum file size in GB",
			},
			&cli.StringFlag{
				Name:  "path",
				Value: ".",
				Usage: "Directory path to scan",
			},
		},
		Action: findLarge,
	}
}

func findLarge(ctx context.Context, c *cli.Command) error {

	path := c.String("path")
	sizeGB := c.Float64("size")

	// Convert GB to bytes
	minBytes := int64(sizeGB * 1024 * 1024 * 1024)

	fmt.Printf("Scanning for files larger than %.2f GB...\n", sizeGB)

	foundAny := false

	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {

		if err != nil {
			fmt.Println("Warning: could not access:", p)
			return nil
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			fmt.Println("Warning: could not stat file:", p)
			return nil
		}

		if info.Size() > minBytes {
			sizeMB := float64(info.Size()) / (1024 * 1024)
			sizeGBActual := sizeMB / 1024
			fmt.Printf("Large file: %s (%.2f GB)\n", p, sizeGBActual)
			foundAny = true
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error while scanning:", err)
	}

	if !foundAny {
		fmt.Printf("No files larger than %.2f GB found.\n", sizeGB)
	}

	return nil
}
