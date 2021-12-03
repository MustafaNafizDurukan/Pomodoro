package main

type paramsT struct {
	PomodoroTime   string   `short:"p" long:"PomodoroTime" required:"false" description:"(Optional) Specifies each pomodoro time. Format: 25m, 1h2m3s" default:"25m"`
	LongBreakTime  string   `short:"l" long:"LongBreakTime" required:"false" description:"(Optional) Specifies long break time. Format: 25m, 1h2m3s" default:"20m"`
	ShortBreakTime string   `short:"s" long:"ShortBreakTime" required:"false" description:"(Optional) Specifies short break time. Format: 25m, 1h2m3s" default:"5m"`
	Config         string   `short:"c" long:"config" required:"false" description:"(Optional) Config file name" default:"config.yml"`
	Args           []string // Positional arguments except application path (1st param) will be here
}

var params paramsT
