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
	CompletedDate time.Time

	P        *pomo.Pomo
	willWait bool
}

// New creates Task structure.
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

// IsCompleted if all pomodoros finished it returns true.
func (t *Task) IsCompleted() bool {
	return !t.CompletedDate.IsZero()
}

// RunOnce runs a pomodoro and if all pomodoros completed returns true
func (t *Task) RunOnce(isPomodoro bool) bool {
	t.assignTaskInfo()

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

// assignTaskInfo assigns all task section, title and message to TaskInfo structure.
func (t *Task) assignTaskInfo() {
	t.P.Timer.TaskInfo.Section = t.Section
	t.P.Timer.TaskInfo.Title = t.Title
	t.P.Timer.TaskInfo.Message = t.Message
}

// Wait waits after pomodoro or break finished.
func (t *Task) Wait(isPomodoro bool) {
	if !t.willWait {
		return
	}

	print.Wait(isPomodoro)

	fmt.Scanln()
}

// AddPomodoro increase pomodoro count.
func (t *Task) AddPomodoro(count int) error {
	if count > 10 {
		return errAddPomodoro
	}

	t.CompletedDate = time.Time{}

	t.P.PomNumber += count

	return nil
}

// SubPomodoro decrease pomodoro count.
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
