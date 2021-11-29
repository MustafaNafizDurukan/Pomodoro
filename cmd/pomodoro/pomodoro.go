package main

import (
	"fmt"
	_ "os"

	"github.com/mustafanafizdurukan/pomodoro/internal/event"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
)

func main() {
	err := console.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// f := font.New("LA", termbox.ColorCyan, termbox.ColorDefault, &font.Position{5, 5})

	// a := termbox.ColorDefault
	// console.Print("sdgd", a, a, 0, 0)

	// f.Echo()
	// console.Flush()

	e, err := event.New("10s")
	if err != nil {
		fmt.Println(err)
		return
	}

	e.Start()
}
