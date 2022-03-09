// Package pomo implements functions that actions about pomodoro.
package pomo

import (
	"errors"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/pkg/logs"
	"github.com/mustafanafizdurukan/pomodoro/pkg/timer"
)

type Pomo struct {
	pomTime            time.Duration
	shortBreak         time.Duration
	PomNumber          int
	CompletedPomNumber int
	CompletedDates     []time.Time

	Timer *timer.Timer
}

var (
	errParse = errors.New("pomo: given time string could not be parsed")
)

// New creates Pomo struct. If given timer string can not be parsed it fails.
func New(pomoTime, shortBreak string, PomNumber int) (*Pomo, error) {
	t, err := timer.New()
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
		pomTime:    pomoTimeD,
		shortBreak: shortBreakD,
		PomNumber:  PomNumber,
		Timer:      t,
	}, err
}

// StartPomodoro starts pomodoro and if all pomodoros completed returns true.
func (p *Pomo) StartPomodoro() bool {
	if p.CompletedPomNumber == p.PomNumber {
		return true
	}

	p.Timer.IsPomodoro = true
	isFinishedNormally := p.Timer.Start(p.pomTime)

	if isFinishedNormally {
		p.CompletedPomNumber++
		p.CompletedDates = append(p.CompletedDates, time.Now())
	}

	p.assignTaskInfo()

	return p.CompletedPomNumber == p.PomNumber
}

// StartBreak starts break.
func (p *Pomo) StartBreak() {
	p.Timer.IsPomodoro = false
	p.Timer.Start(p.shortBreak)
}

// assignTaskInfo assigns all pomodoro count and completed pomodoro count to TaskInfo structure.
func (p *Pomo) assignTaskInfo() {
	p.Timer.TaskInfo.PomNumber = p.PomNumber
	p.Timer.TaskInfo.CompletedPomNumber = p.CompletedPomNumber
}

// RemainingPomodoroCount returns remaining pomodoro count.
func (p *Pomo) RemainingPomodoroCount() int {
	return p.PomNumber - p.CompletedPomNumber
}
