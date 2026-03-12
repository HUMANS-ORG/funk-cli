package commands

import (
	"context"
	"fmt"
	"funk/sqldb"
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli/v3"
	"os/exec"
	"runtime"
	"time"
)

func TimerCommand() *cli.Command {
	return &cli.Command{
		Name:   "timer",
		Usage:  "Set a countdown timer and show Windows toast when done",
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
				Name:  "his",
				Usage: "show the timer history",
			},
		},
	}
}

func TimerSet(ctx context.Context, cmd *cli.Command) error {

	var totalSeconds int

	if cmd.Bool("his") {

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

	fmt.Println("Timer Start")

	switch runtime.GOOS {
	case "windows":
		fmt.Printf("Timer started — %d seconds\n", totalSeconds)
		psCommand := fmt.Sprintf(
			"Start-Sleep -Seconds %d; Import-Module BurntToast; New-BurntToastNotification -Text '%d Seconds Timer Finished'",
			totalSeconds, totalSeconds,
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
		err_d := termbox.Init()

		if err_d != nil {
			panic(err_d)
		}

		defer termbox.Close()

		pause := make(chan bool)

		go func() {
			for {
				ev := termbox.PollEvent()

				if ev.Type == termbox.EventKey {

					if ev.Key == termbox.KeyCtrlC {
						pause <- true
						return
					}

					if ev.Ch =='q'{
						pause <-true
						return 
					}
				}
			}
		}()

		var h int
		var m int
		var s int

		h1, m1, s1 := timer_cal(totalSeconds)

		for i := totalSeconds; i >= 0; i-- {

			h, m, s = timer_cal(i)

			select {
			case <-pause:
				sqldb.Insert_data(h, m, s)
				fmt.Println("\nTimer stopped")
				return nil

			default:
				fmt.Printf("\r⏳ %02d:%02d:%02d", h, m, s)
				time.Sleep(time.Second)

				if i == 0 {
					sqldb.Insert_data(h1, m1, s1)
				}
			}

		}

		fmt.Println("\ntimer finish")
	}

	return nil
}

func timer_cal(i int) (int, int, int) {
	h := i / 3600
	m := (i % 3600) / 60
	s := i % 60
	return h, m, s
}
