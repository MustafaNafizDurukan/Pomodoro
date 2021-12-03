package main

import (
	"fmt"
	"os"
	_ "os"

	"github.com/mustafanafizdurukan/pomodoro/internal/event"
	"github.com/mustafanafizdurukan/pomodoro/internal/pomo"
	"github.com/mustafanafizdurukan/pomodoro/pkg/config"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/flags"
	"github.com/mustafanafizdurukan/pomodoro/pkg/font"
	"github.com/mustafanafizdurukan/pomodoro/pkg/panik"
	"github.com/nsf/termbox-go"
)

func main() {
	var err error
	defer panik.Catch()

	params.Args, err = flags.Parse(&params, os.Args)
	if err != nil {
		fmt.Printf("Error parsing command line arguments: %v \n", err)
		return
	}

	err = config.Init(params.Config)
	if err != nil {
		fmt.Println(nil, "[!] Can not load config file: %s, %v", params.Config, err)
		return
	}

	params.equalizeToConfig()

	err = console.Init()
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

	p, err := pomo.New(params.PomodoroTime, params.ShortBreakTime, params.LongBreakTime, e)
	if err != nil {
		fmt.Println(err)
		return
	}

	p.Start()
}

func (p *paramsT) equalizeToConfig() {
	p.LongBreakTime = config.Config.LongBreakTime
	p.PomodoroTime = config.Config.PomodoroTime
	p.ShortBreakTime = config.Config.ShortBreakTime
}
