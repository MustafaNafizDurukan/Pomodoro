// Package convert implements functions to convert variable types to string
package convert

import (
	"errors"
	"fmt"
	"time"
)

var (
	errParse = errors.New("convert: given time string could not be parsed")
)

// DateToString converts date to string and returns that string
func DateToString(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	if h < 1 {
		return fmt.Sprintf("%02d:%02d", m, s)
	}
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// StringToDate converts string to date and returns that date
func StringToDate(d string) (time.Duration, error) {
	time, err := time.ParseDuration(d)
	if err != nil {
		return 0, errParse
	}

	return time, nil
}
