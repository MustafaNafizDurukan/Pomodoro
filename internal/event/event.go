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
	"github.com/mustafanafizdurukan/pomodoro/pkg/play"
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
func (e *Event) Start(willWait bool) {
	print.Time(e.f, e.TimeLeft)

	if willWait {
		fmt.Scanln()
	}

	e.start()
}

func (e *Event) start() {
	wilRun := true
	defer func() {
		console.Clear()
		console.Flush()
	}()

	timer.Start(e.TimeLeft)

loop:
	for {
		select {
		case ev := <-e.queues:
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				d, _ := convert.StringToDate("10s")
				if e.count(d, print.Quit) {
					os.Exit(1)
				}
			}
			if ev.Ch == 'p' || ev.Ch == 'P' {
				timer.Stop()
			}
			if ev.Ch == 'c' || ev.Ch == 'C' {
				timer.Start(e.TimeLeft)
			}
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeySpace) {
				break loop
			}
		case <-timer.Ticker.C:
			termbox.Sync()
			timer.Decrease(&e.TimeLeft)
			if wilRun {
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

	go play.Sound()
	print.Zero(e.f)
}

type callBack func(d time.Duration)

func (e *Event) count(d time.Duration, cb callBack) bool {
	defer func() {
		console.Clear()
		console.Flush()
	}()

	wilRun := true

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
				cb(d)

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

	return false
}
