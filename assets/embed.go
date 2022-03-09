package assets

import (
	"embed"
	"fmt"
	"io/fs"
)

type Warning string

const (
	StoppedPomodoroSource Warning = "sounds/stopped_pomodoro_session.mp3"
	StoppedBreakSource    Warning = "sounds/you_took_long_break.mp3"
)

//go:embed sounds/*
var f embed.FS

func LoadEmbeddedWarning(warning Warning) (fs.File, error) {
	fs, err := f.Open(string(warning))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return fs, nil
}
