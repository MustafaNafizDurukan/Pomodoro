package pomo

import (
	"fmt"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/print"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/convert"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/mustafanafizdurukan/pomodoro/pkg/logs"
	"github.com/mustafanafizdurukan/pomodoro/pkg/play"
	"github.com/mustafanafizdurukan/pomodoro/pkg/ticker"
	"github.com/nsf/termbox-go"
)

type Timer struct {
	queues chan termbox.Event
	f      *font.Font
}

// New returns pointer of event structure. If given string could not be parsed It returns error.
// Time string should be 1h6m, 23m3s format.
func new(f *font.Font) (*Timer, error) {
	queues := make(chan termbox.Event)
	go func() {
		for {
			queues <- termbox.PollEvent()
		}
	}()

	return &Timer{
		queues: queues,
		f:      f,
	}, nil
}

func (t *Timer) Wait(willWait bool, timeLeft time.Duration) {
	print.Time(t.f, timeLeft)

	if willWait {
		fmt.Scanln()
	}
}

// Start starts pomodoro timer and if pomodoro finishes normally it returns true.
func (t *Timer) Start(timeLeft time.Duration) bool {
	return t.start(timeLeft)
}

func (t *Timer) start(timeLeft time.Duration) bool {
	var err error
	wilRun := true
	defer func() {
		console.Clear()
		console.Flush()
	}()

	ticker.Start(timeLeft)

loop:
	for {
		select {
		case ev := <-t.queues:
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				d, _ := convert.StringToDate("10s")
				if t.count(d, print.Quit) {
					return false
					// os.Exit(1)
				}
			}
			if ev.Ch == 'p' || ev.Ch == 'P' {
				ticker.Stop()
			}
			if ev.Ch == 'c' || ev.Ch == 'C' {
				ticker.Start(timeLeft)
			}
			// if ev.Type == termbox.EventKey && (ev.Key == termbox.KeySpace) {
			// 	break loop
			// }
		case <-ticker.Ticker.C:
			termbox.Sync()
			ticker.Decrease(&timeLeft)
			if wilRun {
				print.Time(t.f, timeLeft)

				wilRun = false
				break
			}
			wilRun = true
		case <-ticker.Timer.C:
			console.Clear()
			break loop
		}
	}

	go func() {
		err = play.Sound()
		if err != nil {
			logs.ERROR.Println(err)
		}
	}()

	print.Zero(t.f)

	return true
}

// type callBack func(d time.Duration)

func (t *Timer) count(d time.Duration, callBack func(d time.Duration)) bool {
	defer func() {
		console.Clear()
		console.Flush()
	}()

	wilRun := true

	ticker.Start(d)

loop:
	for {
		select {
		case ev := <-t.queues:
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				return true
			}
			if ev.Ch == 'n' || ev.Ch == 'N' {
				return false
			}
			if ev.Ch == 'y' || ev.Ch == 'Y' {
				return true
			}
		case <-ticker.Ticker.C:
			termbox.Sync()

			if wilRun {
				callBack(d)

				ticker.Decrease(&d)
				wilRun = false
				break
			}
			ticker.Decrease(&d)
			wilRun = true
		case <-ticker.Timer.C:
			console.Clear()
			break loop
		}
	}

	return false
}
