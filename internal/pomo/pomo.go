// Package pomo implements functions that actions about pomodoro
package pomo

import (
	"errors"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/mustafanafizdurukan/pomodoro/pkg/logs"
	"github.com/nsf/termbox-go"
)

type Pomo struct {
	time               time.Duration
	shortBreak         time.Duration
	PomNumber          int
	CompletedPomNumber int
	CompletedDates     []time.Time

	t *Timer
}

var (
	errParse = errors.New("pomo: given time string could not be parsed")
)

// New creates new Pomo struct. If given timer string can not be parsed it fails.
func New(pomoTime, shortBreak string, PomNumber int) (*Pomo, error) {
	_, y := console.SizeSixteenOver(8)
	x, _ := console.MidPoint()
	pos := font.Position{x, y}

	f := font.New(termbox.ColorCyan, termbox.ColorDefault, &pos)

	t, err := new(f)
	if err != nil {
		logs.ERROR.Println(err)
		return nil, err
	}

	pomoTimeD, err := time.ParseDuration(pomoTime)
	if err != nil {
		return nil, errParse
	}

	shortBreakD, err := time.ParseDuration(shortBreak)
	if err != nil {
		return nil, errParse
	}

	return &Pomo{
		time:       pomoTimeD,
		shortBreak: shortBreakD,
		PomNumber:  PomNumber,
		t:          t,
	}, err
}

// Start starts pomodoro and if all pomodoros completed returns true
func (p *Pomo) StartPomodoro() bool {
	if p.CompletedPomNumber == p.PomNumber {
		return true
	}

	isFinishedNormally := p.t.Start(p.time)

	if isFinishedNormally {
		p.CompletedPomNumber++
		p.CompletedDates = append(p.CompletedDates, time.Now())
	}

	return p.CompletedPomNumber == p.PomNumber
}

func (p *Pomo) StartBreak() {
	p.t.Start(p.shortBreak)
}

func (p *Pomo) RemainingPomodoroCount() int {
	return p.PomNumber - p.CompletedPomNumber
}
