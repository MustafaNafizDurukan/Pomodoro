// Package event implements functions to
// TODO: Complete here
package event

import (
	"errors"
	"os"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/convert"
	"github.com/mustafanafizdurukan/pomodoro/pkg/timer"
	"github.com/nsf/termbox-go"
)

var (
	errParse = errors.New("event: given time string could not be parsed")
)

type Event struct {
	timeLeft time.Duration
	queues   chan termbox.Event
}

// New returns pointer of event structure. If given string could not be parsed It returns error.
// Time string should be 1h6m, 23m3s format.
func New(timeString string) (*Event, error) {
	duration, err := time.ParseDuration(timeString)
	if err != nil {
		return nil, errParse
	}

	queues := make(chan termbox.Event)
	go func() {
		for {
			queues <- termbox.PollEvent()
		}
	}()

	return &Event{
		timeLeft: duration,
		queues:   queues,
	}, nil
}

func (e *Event) Start() {
	wilRun := true
	timer.Start(e.timeLeft)

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
				timer.Start(e.timeLeft)
			}
		case <-timer.Ticker.C:
			timer.Decrease(&e.timeLeft)
			if wilRun {
				console.Print(convert.DateToString(e.timeLeft), termbox.ColorDefault, termbox.ColorDefault, 5, 5)
				console.Flush()
				wilRun = false
				break
			}
			wilRun = true
		case <-timer.Timer.C:
			break loop
		}
	}
}
