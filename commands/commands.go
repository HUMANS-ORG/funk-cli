package commands

import (
	"context"
	"fmt"
	"time"
	"github.com/urfave/cli/v3"
)

func TimerCommand() *cli.Command {
	return  &cli.Command{
		Name: "timer",
		Usage: "detect a particular timer",
		Action: TimerSet,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "seconds",
				Usage: "the timer in seconds",
			},
			&cli.IntFlag{
				Name: "minutes",
				Usage: "the timer in minutes",
			},
			&cli.IntFlag{
				Name: "hours",
				Usage: "the timer in hours",
			},
		},
	}
}

func TimerSet(ctx context.Context,cmd *cli.Command) error  {
	var totalSeconds int

	if cmd.IsSet("seconds") {
		totalSeconds = cmd.Int("seconds")

	} else if cmd.IsSet("minutes") {
		totalSeconds = cmd.Int("minutes") * 60

	} else if cmd.IsSet("hours") {
		totalSeconds = cmd.Int("hours") * 3600

	} else {
		return fmt.Errorf("please specify --seconds, --minutes, or --hours")
	}

    fmt.Println("⏳ Timer started")

	for i := totalSeconds; i > 0; i-- {
		fmt.Printf("Remaining: %d seconds\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("🔔 Time's up!")

	return  nil

}