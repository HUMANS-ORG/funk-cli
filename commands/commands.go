package commands

import (
	"context"
	"fmt"
	"os/exec"
	"github.com/urfave/cli/v3"
)

func TimerCommand() *cli.Command {
	return &cli.Command{
		Name:  "timer",
		Usage: "Set a countdown timer and show Windows toast when done",
		Action: TimerSet,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "seconds",
				Usage: "timer duration in seconds",
			},
			&cli.IntFlag{
				Name:  "minutes",
				Usage: "timer duration in minutes",
			},
			&cli.IntFlag{
				Name:  "hours",
				Usage: "timer duration in hours",
			},
		},
	}
}

func TimerSet(ctx context.Context, cmd *cli.Command) error {

	var totalSeconds int


	switch {
	case cmd.IsSet("seconds"):
		totalSeconds = cmd.Int("seconds")

	case cmd.IsSet("minutes"):
		totalSeconds = cmd.Int("minutes") * 60

	case cmd.IsSet("hours"):
		totalSeconds = cmd.Int("hours") * 3600

	default:
		return fmt.Errorf("please specify one of: --seconds, --minutes, or --hours")
	}

	if totalSeconds <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	fmt.Printf("⏳ Timer started — %d seconds\n", totalSeconds)

	// PowerShell background timer command
	psCommand := fmt.Sprintf(
		"Start-Sleep -Seconds %d; Import-Module BurntToast; New-BurntToastNotification -Text '%d Seconds Timer Finished'",
		totalSeconds,totalSeconds,
	)

	// run powershell in background
	cmdExec := exec.Command(
		"powershell.exe",
		"-NoProfile",
		"-Command",
		psCommand,
	)

	err := cmdExec.Start()
	if err != nil {
		return fmt.Errorf("failed to start timer: %v", err)
	}

	fmt.Println("✅ Timer running in background")

	return nil
}