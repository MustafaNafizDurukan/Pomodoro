// Package pomo implements functions that actions about pomodoro
package pomo

import (
	"errors"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/event"
)

type Pomo struct {
	pomoTime   time.Duration
	shortBreak time.Duration
	longBreak  time.Duration
	e          *event.Event
}

var (
	errParse = errors.New("pomo: given time string could not be parsed")
)

// New creates new Pomo struct. If given timer string can not be parsed it fails.
func New(pomoTime, shortBreak, longBreak string, e *event.Event) (*Pomo, error) {
	pomoTimeD, err := time.ParseDuration(pomoTime)
	if err != nil {
		return nil, errParse
	}

	shortBreakD, err := time.ParseDuration(shortBreak)
	if err != nil {
		return nil, errParse
	}

	longBreakD, err := time.ParseDuration(longBreak)
	if err != nil {
		return nil, errParse
	}

	return &Pomo{
		pomoTime:   pomoTimeD,
		shortBreak: shortBreakD,
		longBreak:  longBreakD,
		e:          e,
	}, err
}

// Start starts pomodoro
func (p *Pomo) Start() error {
	for i := 0; i < 12; i++ {
		p.e.PomoCount = i
		if i != 0 && i%4 == 0 {
			p.e.TimeLeft = p.longBreak
			p.e.Start()
		}

		p.e.TimeLeft = p.pomoTime
		p.e.Start()

		p.e.TimeLeft = p.shortBreak
		p.e.Start()
	}

	return nil
}
