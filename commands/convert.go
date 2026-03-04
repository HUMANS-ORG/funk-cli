package commands

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
)

func ConvertCommand() *cli.Command {
	return &cli.Command{
		Name:   "convert",
		Usage:  "converts various units",
		Action: Converto,
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "km",
				Usage: "Convert kilometers to Miles",
			},
			&cli.Float64Flag{
				Name:  "miles",
				Usage: "Convert miles to kilometers",
			},
		},
	}
}

func Converto(ctx context.Context, cmd *cli.Command) error {
	if cmd.IsSet("km") {
		km := cmd.Float64("km")
		miles := km * 0.6213712
		fmt.Printf("%.2f km = %.2f miles \n", km, miles)
		return nil
	} else if cmd.IsSet("miles") {
		miles := cmd.Float64("miles")
		km := miles / 0.6213712
		fmt.Printf("%.2f miles = %.2f km \n", miles, km)
		return nil
	} else {
		return fmt.Errorf("Please specify either --km or --miles")
	}

}
