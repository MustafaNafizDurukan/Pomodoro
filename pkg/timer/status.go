package timer

type Status string

const (
	running  Status = "timer-running"
	stopped  Status = "timer-stopped"
	finished Status = "timer-finished"
)
