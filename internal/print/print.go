package print

import (
	"fmt"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/convert"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/nsf/termbox-go"
)

var (
	pomoMotivMsg string
	ShouldAlign  bool
)

type TaskInfo struct {
	Section string
	Title   string
	Message string

	PomNumber          int
	CompletedPomNumber int
}

// Time prints left time, message and task info to the console.
func Time(f *font.Font, TimeLeft time.Duration, ti *TaskInfo) {
	f.Text = convert.DateToString(TimeLeft)

	if f.Text == "" {
		return
	}

	console.Clear()
	defer console.Flush()

	printMotivationMessage(TimeLeft)

	printTaskInfo(ti)

	calculateFontPoints(f)
	f.Echo()
}

var (
	sectionMsg = "Section: %s"
	titleMsg   = "Title: %s"
	pomMsg     = "[%d/%d] Pomodoro completed"
)

// printTaskInfo prints task info to the console.
func printTaskInfo(ti *TaskInfo) {
	sectionMessage := fmt.Sprintf(sectionMsg, ti.Section)
	titleMessage := fmt.Sprintf(titleMsg, ti.Title)
	pomMessage := fmt.Sprintf(pomMsg, ti.CompletedPomNumber, ti.PomNumber)

	x, y := console.SizeSixteenOver(1)
	console.Print(sectionMessage, termbox.ColorDefault, termbox.ColorDefault, x, y)
	console.Print(titleMessage, termbox.ColorDefault, termbox.ColorDefault, x, y+2)
	console.Print(pomMessage, termbox.ColorDefault, termbox.ColorDefault, x, y+4)
}

// printMotivationMessage prints motivation message to the console.
func printMotivationMessage(TimeLeft time.Duration) {
	m := TimeLeft.Round(time.Second)
	if int(m.Seconds())%10 == 0 {
		pomoMotivMsg = message()
	}

	_, y := console.SizeSixteenOver(11)
	x, _ := console.MidPoints()
	console.Print(pomoMotivMsg, termbox.ColorDefault, termbox.ColorDefault, x-len(pomoMotivMsg)/2, y)
}

// calculateFontPoints calculates font position by calculating font size.
func calculateFontPoints(f *font.Font) {
	fo := font.ToFont(f.Text)

	x, y := console.MidPoints()
	w, h := font.Size(fo)

	f.Position.X = x - w/2
	f.Position.Y = y - h/2
}

// Zero prints 00:00 for 2 seconds to the console
func Zero(f *font.Font) {
	f.Text = "00:00"

	for i := 0; i < 3; i++ {
		console.Clear()
		console.Flush()
		time.Sleep(time.Second / 2)

		f.Echo()
		console.Flush()
		time.Sleep(time.Second / 2)
	}
}

// Quit prints quit message when you press q on keyboard
func Quit(d time.Duration) {
	console.Clear()
	x, y := console.MidPoints()
	msg := fmt.Sprintf("Are you sure want to quit? (No:n, Yes:y) %s", d.String())
	console.Print(msg, termbox.ColorDefault, termbox.ColorDefault, x-len(msg)/2, y)
	console.Flush()

	msg = "Current session will be lost."
	console.Print(msg, termbox.ColorDefault, termbox.ColorDefault, x-len(msg)/2, y+2)
	console.Flush()
}

// Wait prints wait message apter pomodoro or break finished
func Wait(isPomodoro bool) {
	console.Clear()
	defer console.Flush()

	x, y := console.MidPoints()

	str := "the break"
	if isPomodoro {
		str = "pomodoro"
	}

	msg := fmt.Sprintf("Press any key once to start %s", str)
	console.Print(msg, termbox.ColorDefault, termbox.ColorDefault, x-len(msg)/2, y)
}
