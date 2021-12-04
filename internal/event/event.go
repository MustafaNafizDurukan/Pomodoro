// Package event implements functions to
// TODO(mustafa): Complete here
package event

import (
	"errors"
	"os"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/print"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
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

func (e *Event) Start() {
	wilRun := true
	timer.Start(e.TimeLeft)

loop:
	for {
		select {
		case ev := <-e.queues:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeySpace) {
				break loop
			}
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyCtrlC) {
				os.Exit(0)
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
