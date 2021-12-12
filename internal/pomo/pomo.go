// Package pomo implements functions that actions about pomodoro
package pomo

import (
	"errors"
	"fmt"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/event"
)

type Pomo struct {
	time       time.Duration
	shortBreak time.Duration
	longBreak  time.Duration
	willWait   bool
	e          *event.Event
}

var (
	errParse = errors.New("pomo: given time string could not be parsed")
)

// New creates new Pomo struct. If given timer string can not be parsed it fails.
func New(pomoTime, shortBreak, longBreak string, WillWait bool, e *event.Event) (*Pomo, error) {
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
		time:       pomoTimeD,
		shortBreak: shortBreakD,
		longBreak:  longBreakD,
		willWait:   WillWait,
		e:          e,
	}, err
}

// Start starts pomodoro
func (p *Pomo) Start() error {
	for i := 0; i < 10240; i++ {
		if p.willWait {
			fmt.Scanln()
		}

		if i%2 == 0 {
			p.e.TimeLeft = p.time
			p.e.Start()
			continue
		}

		if i%8 == 7 {
			p.e.TimeLeft = p.longBreak
			p.e.Start()
			continue
		}

		p.e.TimeLeft = p.shortBreak
		p.e.Start()
	}

	return nil
}
