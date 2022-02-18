package timer

import (
	"time"

	"github.com/mustafanafizdurukan/pomodoro/internal/print"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/convert"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/mustafanafizdurukan/pomodoro/pkg/logs"
	"github.com/mustafanafizdurukan/pomodoro/pkg/play"
	"github.com/nsf/termbox-go"
)

type Timer struct {
	queues   chan termbox.Event
	f        *font.Font
	TaskInfo *print.TaskInfo
}

// New returns pointer of event structure. If given string could not be parsed It returns error.
// Time string should be 1h6m, 23m3s format.
func New() (*Timer, error) {
	_, y := console.SizeSixteenOver(8)
	x, _ := console.MidPoints()
	pos := font.Position{x, y}

	f := font.New(termbox.ColorCyan, termbox.ColorDefault, &pos)

	ti := &print.TaskInfo{}

	queues := make(chan termbox.Event)
	go func() {
		for {
			queues <- termbox.PollEvent()
		}
	}()

	return &Timer{
		queues:   queues,
		f:        f,
		TaskInfo: ti,
	}, nil
}

// Start starts pomodoro timer and if pomodoro finishes normally it returns true.
func (t *Timer) Start(timeLeft time.Duration) bool {
	var err error
	wilRun := true
	defer func() {
		console.Clear()
		console.Flush()
	}()

	start(timeLeft)

loop:
	for {
		select {
		case ev := <-t.queues:
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				d, _ := convert.StringToDate("10s")
				if t.count(d, print.Quit) {
					return false
				}
			}
			if ev.Ch == 'p' || ev.Ch == 'P' {
				stop()
			}
			if ev.Ch == 'c' || ev.Ch == 'C' {
				start(timeLeft)
			}
		case <-tckr.C:
			termbox.Sync()
			decrease(&timeLeft)
			if wilRun {
				print.Time(t.f, timeLeft, t.TaskInfo)

				wilRun = false
				break
			}
			wilRun = true
		case <-tmr.C:
			console.Clear()
			break loop
		}
	}

	go func() {
		err = play.Sound()
		if err != nil {
			logs.ERROR.Println(err)
		}
	}()

	print.Zero(t.f)

	return true
}

// type callBack func(d time.Duration)

func (t *Timer) count(d time.Duration, callBack func(d time.Duration)) bool {
	defer func() {
		console.Clear()
		console.Flush()
	}()

	wilRun := true

	start(d)

loop:
	for {
		select {
		case ev := <-t.queues:
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				return true
			}
			if ev.Ch == 'n' || ev.Ch == 'N' {
				return false
			}
			if ev.Ch == 'y' || ev.Ch == 'Y' {
				return true
			}
		case <-tckr.C:
			termbox.Sync()

			if wilRun {
				callBack(d)

				decrease(&d)
				wilRun = false
				break
			}
			decrease(&d)
			wilRun = true
		case <-tmr.C:
			console.Clear()
			break loop
		}
	}

	return false
}
