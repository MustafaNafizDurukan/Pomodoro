package timer

import (
	"time"

	"github.com/mustafanafizdurukan/pomodoro/assets"
	"github.com/mustafanafizdurukan/pomodoro/pkg/play"
)

var (
	tmr  *time.Timer
	tckr *time.Ticker
)

const (
	tick = time.Second / 2
)

// start starts timer and ticker
func (t *Timer) start(d time.Duration) {
	tmr = time.NewTimer(d)
	tckr = time.NewTicker(tick)

	t.Status = running
}

// stop stops timer and ticker
func (t *Timer) stop() {
	tmr.Stop()
	tckr.Stop()

	t.Status = stopped
}

// decrease decreases left time
func decrease(timeLeft *time.Duration) {
	*timeLeft -= time.Duration(tick)
}

const (
	shortBreak = time.Duration(10 * time.Minute)
	longBreak  = time.Duration(20 * time.Minute)
)

func (t *Timer) startPauseTimer() {
	counter := 0
	for {
		if 10 < counter {
			break
		}

		if counter < 4 {
			time.Sleep(shortBreak)
		} else {
			time.Sleep(longBreak)
		}

		if t.Status == running {
			break
		}

		if t.IsPomodoro {
			play.Warning(assets.StoppedPomodoroSource)
		} else {
			play.Warning(assets.StoppedBreakSource)
		}

		counter++
	}
}
