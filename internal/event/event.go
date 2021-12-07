// Package event implements functions to
// TODO(mustafa): Complete here
package event

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/print"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/convert"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/mustafanafizdurukan/pomodoro/pkg/timer"
	"github.com/nsf/termbox-go"
)

var (
	errParse = errors.New("event: given time string could not be parsed")
)

type Event struct {
	TimeLeft time.Duration
	f        *font.Font
	queues   chan termbox.Event
}

// New returns pointer of event structure. If given string could not be parsed It returns error.
// Time string should be 1h6m, 23m3s format.
func New(f *font.Font) (*Event, error) {
	queues := make(chan termbox.Event)
	go func() {
		for {
			queues <- termbox.PollEvent()
		}
	}()

	return &Event{
		queues: queues,
		f:      f,
	}, nil
}

// Start starts pomodoro timer
func (e *Event) Start() {
	wilRun := true
	timer.Start(e.TimeLeft)

loop:
	for {
		select {
		case ev := <-e.queues:
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				if e.quit() {
					os.Exit(1)
				}
			}
			if ev.Ch == 'p' || ev.Ch == 'P' {
				timer.Stop()
			}
			if ev.Ch == 'c' || ev.Ch == 'C' {
				timer.Start(e.TimeLeft)
			}
		case <-timer.Ticker.C:
			termbox.Sync()
			timer.Decrease(&e.TimeLeft)
			if wilRun {
				console.Clear()
				print.Time(e.f, e.TimeLeft)

				wilRun = false
				break
			}
			wilRun = true
		case <-timer.Timer.C:
			console.Clear()
			break loop
		}
	}

	print.Zero(e.f)
}

func (e *Event) quit() bool {
	console.Clear()
	console.Flush()

	wilRun := true

	d, _ := convert.StringToDate("10s")

	timer.Start(d)

loop:
	for {
		select {
		case ev := <-e.queues:
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				return true
			}
			if ev.Ch == 'n' || ev.Ch == 'N' {
				return false
			}
			if ev.Ch == 'y' || ev.Ch == 'Y' {
				return true
			}
		case <-timer.Ticker.C:
			termbox.Sync()

			if wilRun {
				x, y := console.MidPoint()
				msg := fmt.Sprintf("Are you sure want to quit? (No:n, Yes:y) %s", d.String())
				console.Print(msg, termbox.ColorDefault, termbox.ColorDefault, x-len(msg)/2, y)
				console.Flush()

				msg = "Current session will be lost."
				console.Print(msg, termbox.ColorDefault, termbox.ColorDefault, x-len(msg)/2, y+2)
				console.Flush()

				timer.Decrease(&d)
				wilRun = false
				break
			}
			timer.Decrease(&d)
			wilRun = true
		case <-timer.Timer.C:
			console.Clear()
			break loop
		}
	}

	console.Clear()
	console.Flush()
	return false
}
