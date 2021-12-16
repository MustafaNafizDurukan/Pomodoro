package main

type paramsT struct {
	PomodoroTime   string   `short:"p" long:"pomodoro" required:"false" description:"(Optional) Specifies each pomodoro time. Format: 25m, 1h2m3s" default:"25m"`
	LongBreakTime  string   `short:"l" long:"longbreak" required:"false" description:"(Optional) Specifies long break time. Format: 25m, 1h2m3s" default:"20m"`
	ShortBreakTime string   `short:"s" long:"shortbreak" required:"false" description:"(Optional) Specifies short break time. Format: 25m, 1h2m3s" default:"5m"`
	Config         string   `short:"c" long:"config" required:"false" description:"(Optional) Config file name" default:"config.yml"`
	WillWait       bool     `short:"w" long:"willwait" required:"false" description:"(Optional) When pomodoro or the break is over will the next stage be passed without waiting?"`
	Music          string   `short:"r" long:"music" required:"false" description:"(Optional) When pomodoro finish will the music play randomly? If you give empty string then it will be played randombly." default:"../sounds/ring.mp3"`
	Args           []string // Positional arguments except application path (1st param) will be here
}

var params paramsT
