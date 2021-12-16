// Package pomo implements functions that actions about pomodoro
package pomo

import (
	"errors"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/event"
)

type Pomo struct {
	time       time.Duration
	shortBreak time.Duration
	longBreak  time.Duration
	willWait   bool
	music      string
	e          *event.Event
}

var (
	errParse = errors.New("pomo: given time string could not be parsed")
)

// New creates new Pomo struct. If given timer string can not be parsed it fails.
func New(pomoTime, shortBreak, longBreak string, WillWait bool, music string, e *event.Event) (*Pomo, error) {
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
		music:      music,
		e:          e,
	}, err
}

// Start starts pomodoro
func (p *Pomo) Start() error {
	for i := 0; i < 10240; i++ {
		if i%2 == 0 {
			p.e.TimeLeft = p.time
			p.e.Start(p.willWait, p.music)
			continue
		}

		if i%8 == 7 {
			p.e.TimeLeft = p.longBreak
			p.e.Start(p.willWait, p.music)
			continue
		}

		p.e.TimeLeft = p.shortBreak
		p.e.Start(p.willWait, p.music)
	}

	return nil
}
