package task

import (
	"errors"
	"fmt"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/constants"
	"github.com/mustafanafizdurukan/pomodoro/internal/pomo"
	"github.com/mustafanafizdurukan/pomodoro/internal/print"
)

var (
	errSubPomodoro = errors.New("pomodoro count can not be lower than completed pomodoro count")
	errAddPomodoro = errors.New("you can not add more than 10 pomodoro at a time")
)

type Task struct {
	id            int
	Section       string //(Daily, Job, ...) Her sectionın belirli özelliği olabilecek mesela Daily sectionının taskları her gün otomatik olarak oluşturulacak
	Title         string
	Message       string
	AddedDate     time.Time
	CompletedDate time.Time //01-02-2022 18:06:56

	P        *pomo.Pomo
	willWait bool
}

func New(section, title, message string, willWait bool, pomodoro *pomo.Pomo) *Task {
	task := &Task{
		id:        constants.TaskId,
		Section:   section,
		Title:     title,
		Message:   message,
		AddedDate: time.Now(),
		P:         pomodoro,
		willWait:  willWait,
	}
	constants.TaskId++
	return task
}

func (t *Task) IsCompleted() bool {
	return !t.CompletedDate.IsZero()
}

// RunOnce runs a pomodoro and if all pomodoros completed returns true
func (t *Task) RunOnce(isPomodoro bool) bool {
	if isPomodoro {
		if t.P.StartPomodoro() {
			t.CompletedDate = time.Now()
			return true
		}
	} else {
		t.P.StartBreak()
	}

	return false
}

func (t *Task) Wait(isPomodoro bool) {
	if !t.willWait {
		return
	}

	print.Wait(isPomodoro)

	fmt.Scanln()
}

func (t *Task) AddPomodoro(count int) error {
	if count > 10 {
		return errAddPomodoro
	}

	t.CompletedDate = time.Time{}

	t.P.PomNumber += count

	return nil
}

func (t *Task) SubPomodoro(count int) error {
	dummy := t.P.PomNumber - count

	if t.P.CompletedPomNumber > dummy {
		return errSubPomodoro
	}

	if t.P.CompletedPomNumber == dummy {
		t.CompletedDate = time.Now()
	}

	t.P.PomNumber = dummy

	return nil
}

/*
type homeContext struct {
	Task    *task.Task
	DNS	    *proxy.DNS
}

type DNS struct {
	Cache            *cache.Cache
	IsBlockingActive bool
	blockedServices  []string
}

type Task struct {
	Section	           string	//(Daily, Job, ...) Her sectionın belirli özelliği olabilecek mesela Daily sectionının taskları her gün otomatik olarak oluşturulacak
	Title              string
	Message            string
	AddedDate 		   time.Duration
	CompletedDate 	   time.Duration    //01-02-2022 18:06:56

	Pomodoro           *pomo.Pomodoro
}

type Pomo struct {
	time               time.Duration
	shortBreak         time.Duration
	PomNumber          int
	CompletedPomNumber int
	CompletedDates 	   []time.Duration

	willWait           bool
	t                  *Timer
}

type Timer struct {
	TimeLeft time.Duration
	queues   chan termbox.Event
	f        *font.Font
}
*/
