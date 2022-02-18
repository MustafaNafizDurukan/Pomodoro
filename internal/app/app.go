// Package app is actually whole program.
package app

import (
	"sync"

	"github.com/mustafanafizdurukan/pomodoro/internal/task"
	"github.com/mustafanafizdurukan/pomodoro/pkg/dnsfilter/proxy"
)

type App struct {
	Tasks       []*task.Task
	ActiveTask  *task.Task
	DNS         *proxy.DNS
	eventsMutex sync.Mutex
	listeners   []EventListener
}

// New creates app structure.
func New(at *task.Task, dns *proxy.DNS) *App {
	return &App{
		ActiveTask: at,
		DNS:        dns,
	}
}

// Add appends new task to task slice on app struct.
func (c *App) Add(t *task.Task) {
	c.Tasks = append(c.Tasks, t)
}

// Run runs for the set number of pomodoros.
func (a *App) Run() {
	for a.ActiveTask.P.PomNumber > a.ActiveTask.P.CompletedPomNumber {
		a.RunOnce()
	}
}

var (
	isPomodoro = 1
)

// RunOnce runs a pomodoro or a break.
func (a *App) RunOnce() {
	if isPomodoro == 1 {
		a.onTaskStarted()
	}
	a.Wait()
	a.ActiveTask.RunOnce(isPomodoro == 1)
	a.onTaskCompleted()

	isPomodoro *= -1
}

// Wait waits after pomodoro or break finished.
func (a *App) Wait() {
	a.ActiveTask.Wait(isPomodoro == 1)
}

// AddEventListener adds new listener to listener slice on app struct.
func (c *App) AddEventListener(listener EventListener) {
	c.listeners = append(c.listeners, listener)
}

// fireEvent runs all listeners on app struct.
func (c *App) fireEvent(e *Event) {
	c.eventsMutex.Lock()
	defer c.eventsMutex.Unlock()
	for _, listener := range c.listeners {
		listener(e)
	}
}
