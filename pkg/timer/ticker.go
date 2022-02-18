package timer

import "time"

var (
	tmr  *time.Timer
	tckr *time.Ticker
)

const (
	tick = time.Second / 2
)

// start starts timer and ticker
func start(d time.Duration) {
	tmr = time.NewTimer(d)
	tckr = time.NewTicker(tick)
}

// stop stops timer and ticker
func stop() {
	tmr.Stop()
	tckr.Stop()
}

// decrease decreases left time
func decrease(timeLeft *time.Duration) {
	*timeLeft -= time.Duration(tick)
}
