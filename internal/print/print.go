package print

import (
	"fmt"
	"time"

	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/convert"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/mustafanafizdurukan/pomodoro/pkg/list"
	"github.com/nsf/termbox-go"
)

var (
	pomoC       string
	ShouldAlign bool
)

// Time prints left time and message to the console
func Time(f *font.Font, TimeLeft time.Duration) {
	f.Text = convert.DateToString(TimeLeft)

	if f.Text == "" {
		return
	}

	console.Clear()
	defer console.Flush()

	m := TimeLeft.Round(time.Second)
	if int(m.Seconds())%10 == 0 {
		pomoC = list.Message()
	}

	_, y := console.SizeSixteenOver(11)
	x, _ := console.MidPoint()

	console.Print(pomoC, termbox.ColorDefault, termbox.ColorDefault, x-len(pomoC)/2, y)

	calculatePoints(f)
	f.Echo()
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

func calculatePoints(f *font.Font) {
	fo := font.ToFont(f.Text)

	x, y := console.MidPoint()
	w, h := font.Size(fo)

	f.Position.X = x - w/2
	f.Position.Y = y - h/2
}

// Quit prints quit message when you press q on keyboard
func Quit(d time.Duration) {
	console.Clear()
	x, y := console.MidPoint()
	msg := fmt.Sprintf("Are you sure want to quit? (No:n, Yes:y) %s", d.String())
	console.Print(msg, termbox.ColorDefault, termbox.ColorDefault, x-len(msg)/2, y)
	console.Flush()

	msg = "Current session will be lost."
	console.Print(msg, termbox.ColorDefault, termbox.ColorDefault, x-len(msg)/2, y+2)
	console.Flush()
}
