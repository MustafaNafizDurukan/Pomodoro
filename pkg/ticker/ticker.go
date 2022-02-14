package ticker

import "time"

var (
	Timer  *time.Timer
	Ticker *time.Ticker
)

const (
	tick = time.Second / 2
)

// Start starts timer and ticker
func Start(d time.Duration) {
	Timer = time.NewTimer(d)
	Ticker = time.NewTicker(tick)
}

// Stop stops timer and ticker
func Stop() {
	Timer.Stop()
	Ticker.Stop()
}

func Decrease(timeLeft *time.Duration) {
	*timeLeft -= time.Duration(tick)
}
