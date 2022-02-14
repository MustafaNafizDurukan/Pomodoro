package main

import (
	"log"
	"os"
	_ "os"

	"github.com/mustafanafizdurukan/pomodoro/internal/home"
	"github.com/mustafanafizdurukan/pomodoro/internal/pomo"
	"github.com/mustafanafizdurukan/pomodoro/internal/task"
	"github.com/mustafanafizdurukan/pomodoro/pkg/config"
	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/mustafanafizdurukan/pomodoro/pkg/dnsfilter/proxy"
	"github.com/mustafanafizdurukan/pomodoro/pkg/flags"
	"github.com/mustafanafizdurukan/pomodoro/pkg/logs"
	"github.com/mustafanafizdurukan/pomodoro/pkg/panik"
)

func main() {
	var err error
	defer panik.Catch()

	// time.Sleep(13 * time.Second)

	log.SetOutput(logs.Set())

	params.Args, err = flags.Parse(&params, os.Args)
	if err != nil {
		logs.ERROR.Printf("Error parsing command line arguments: %v \n", err)
		return
	}

	err = config.Init(params.Config)
	if err != nil {
		logs.ERROR.Printf("[!] Can not load config file: %s, %v", params.Config, err)
		return
	}

	params.equalizeToConfig()

	err = console.Init()
	if err != nil {
		logs.ERROR.Println(err)
		return
	}

	p, err := pomo.New(
		params.PomodoroTime,
		params.ShortBreakTime,
		10,
	)
	if err != nil {
		logs.ERROR.Println(err)
		return
	}

	t := task.New("Daily", "Title", "message", params.WillWait, p)

	DnsProxy := proxy.New()

	DnsProxy.InitBlockedServices("youtube", "twitter")
	DnsProxy.IsBlockingActive = false

	DnsProxy.Start()

	app := home.New(t, DnsProxy)

	app.AddEventListener(func(e *home.Event) {
		app.DNS.IsBlockingActive = e.IsProxyEnabled
	})

	app.Run()
}

func (p *paramsT) equalizeToConfig() {
	p.PomodoroTime = config.Config.PomodoroTime
	p.ShortBreakTime = config.Config.ShortBreakTime
	p.WillWait = config.Config.WillWait
}
