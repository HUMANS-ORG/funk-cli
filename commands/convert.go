package commands

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
)

func ConvertCommand() *cli.Command {
	return &cli.Command{
		Name:    "convert",
		Suggest: true,
		Usage:   "converts various units",
		Action:  Converto, Flags: []cli.Flag{
			// distance converters.
			&cli.Float64Flag{
				Name:    "miles",
				Aliases: []string{"m"},
				Usage:   "Enter values in miles",
			},
			&cli.Float64Flag{
				Name:    "km",
				Aliases: []string{"k"},
				Usage:   "Enter values in kilometers",
			},
			&cli.BoolFlag{
				Name:    "to-km",
				Aliases: []string{"tk"},
				Usage:   "convert to kilometer",
			},
			&cli.BoolFlag{
				Name:    "to-miles",
				Aliases: []string{"tM"},
				Usage:   "Convert to miles",
			},
			&cli.BoolFlag{
				Name:    "to-meters",
				Aliases: []string{"tm"},
				Usage:   "convert to meters",
			},
			// Weight converters
			&cli.Float64Flag{
				Name:    "lbs",
				Aliases: []string{"p"},
				Usage:   "Convert kilograms to pounds",
			},
			&cli.Float64Flag{
				Name:    "kg",
				Aliases: []string{"w"},
				Usage:   "Convert pounds to kilograms",
			},
		},
		//	OnUsageError: ErrorHandle,
	}
}

func Converto(ctx context.Context, cmd *cli.Command) error {

	// converts kilometer values to miles and vice versa
	km := cmd.Float64("km")
	if cmd.IsSet("km") {
		if cmd.Bool("to-miles") {
			miles := km * 0.621371
			fmt.Printf("\n%.2f km = %.2f miles \n", km, miles)
		}
		if cmd.Bool("to-meters") {
			meters := km * 1000
			fmt.Printf("\n%.2f km = %.2f meters\n", km, meters)
		}
	}

	miles := cmd.Float64("miles")
	if cmd.IsSet("miles") {
		if cmd.Bool("to-km") {
			kim := miles / 0.621371
			fmt.Printf("\n%.2f miles = %.2f km \n", miles, kim)
		}
		if cmd.Bool("to-meters") {
			meters := miles * 1609.344
			fmt.Printf("\n%.2f miles = %.2f meters \n", miles, meters)
		}
	}

	// Converts kilograms to pounds and vice versa

	return nil

	// Converts Celsius to Fahrenheit

}

func ErrorHandle(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
	return nil
}
