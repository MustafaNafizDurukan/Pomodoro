package main

import (
	"fmt"
	_ "os"

	"github.com/mustafanafizdurukan/pomodoro/internal/event"
	"github.com/mustafanafizdurukan/pomodoro/internal/pomo"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/nsf/termbox-go"
)

func main() {
	err := console.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, y := console.SizeSixteenOver(6)
	x, _ := console.MidPoint()
	pos := font.Position{x, y}

	f := font.New(termbox.ColorCyan, termbox.ColorDefault, &pos)

	e, err := event.New(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	p, err := pomo.New("8s", "1s", "2s", e)
	if err != nil {
		fmt.Println(err)
		return
	}

	p.Start()
}
