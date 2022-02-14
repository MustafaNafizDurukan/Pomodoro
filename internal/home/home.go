package home

import (
	"sync"

	"github.com/mustafanafizdurukan/pomodoro/internal/task"
	"github.com/mustafanafizdurukan/pomodoro/pkg/dnsfilter/proxy"
)

type app struct {
	Tasks       []*task.Task
	ActiveTask  *task.Task
	DNS         *proxy.DNS
	eventsMutex sync.Mutex
	listeners   []EventListener
}

func New(at *task.Task, dns *proxy.DNS) *app {
	return &app{
		ActiveTask: at,
		DNS:        dns,
	}
}

func (c *app) Add(t *task.Task) {
	c.Tasks = append(c.Tasks, t)
}

func (a *app) Run() {
	for a.ActiveTask.P.PomNumber > a.ActiveTask.P.CompletedPomNumber {
		a.RunOnce()
	}
}

var (
	isPomodoro = 1
)

func (a *app) RunOnce() {
	if isPomodoro == 1 {
		a.onTaskStarted()
	}
	a.Wait()
	a.ActiveTask.RunOnce(isPomodoro == 1)
	a.onTaskCompleted()

	isPomodoro *= -1
}

func (a *app) Wait() {
	a.ActiveTask.Wait(isPomodoro == 1)
}

func (c *app) AddEventListener(listener EventListener) {
	c.listeners = append(c.listeners, listener)
}

func (c *app) fireEvent(e *Event) {
	c.eventsMutex.Lock()
	defer c.eventsMutex.Unlock()
	for _, listener := range c.listeners {
		listener(e)
	}
}
