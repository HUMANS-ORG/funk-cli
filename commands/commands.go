package commands

import (
	"context"
	"fmt"
	"os/exec"
	"github.com/urfave/cli/v3"
	"runtime"
	"time"
	"funk/sqldb"
)

func TimerCommand() *cli.Command {
	return &cli.Command{
		Name:  "timer",
		Usage: "Set a countdown timer and show Windows toast when done",
		Action: TimerSet,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "sec",
				Usage: "timer duration in seconds",
			},
			&cli.IntFlag{
				Name:  "min",
				Usage: "timer duration in minutes",
			},
			&cli.IntFlag{
				Name:  "hr",
				Usage: "timer duration in hours",
			},
			&cli.BoolFlag{
				Name: "his",
				Usage: "show the timer history",
			},
		},
	}
}

func TimerSet(ctx context.Context, cmd *cli.Command) error {

	var totalSeconds int

	if cmd.Bool("his"){
	
		sqldb.Show_history()
		return nil
	}

	switch {
	case cmd.IsSet("sec"):
		totalSeconds = cmd.Int("sec")

	case cmd.IsSet("min"):
		totalSeconds = cmd.Int("min") * 60

	case cmd.IsSet("hr"):
		totalSeconds = cmd.Int("hr") * 3600

	default:
		return fmt.Errorf("please specify one of: --sec, --min, or --hr")
	}

	if totalSeconds <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	sqldb.Create_db()
	
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("Timer started — %d seconds\n", totalSeconds)
		psCommand := fmt.Sprintf(
		"Start-Sleep -Seconds %d; Import-Module BurntToast; New-BurntToastNotification -Text '%d Seconds Timer Finished'",
		totalSeconds,totalSeconds,
	)

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

	fmt.Println("Timer running in background")

	default:
		var h int
		var m int
		var s int

		h,m,s,totalSeconds := timer_cal(h,m,s,totalSeconds)

		sqldb.Insert_data(h,m,s)
		
		for i := totalSeconds ;i>=0;i-- {
			h,m,s,i=timer_cal(h,m,s,i)

			fmt.Printf("\r⏳ %02d:%02d:%02d", h,m,s)

			time.Sleep(time.Second)

		}

		fmt.Println("\ntimer finish")
	}
	
	return nil
}

func timer_cal(h int,m int,s int,i int) (int,int,int,int) {
			h =i/3600
			m = (i%3600)/60
			s =i % 60
			return h,m,s,i
		}