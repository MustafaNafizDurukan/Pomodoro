// Package console implement functions to action about console
package console

import (
	"errors"
	"math"

	"github.com/nsf/termbox-go"
)

var (
	errInit  = errors.New("termbox: could not be initialized")
	errClear = errors.New("termbox: could not be cleared")
	errFlush = errors.New("termbox: could not be flushed")
)

// Size return consoles width and height.
func Size() (int, int) {
	w, h := termbox.Size()
	return w, h
}

// MidPoint returns middle point of console
func MidPoints() (int, int) {
	w, h := Size()
	return w / 2, h / 2
}

// SizeSixteenOver returns returns Number/16 of the
// width and height of the console based on the number you enter.
// If width of height of console lower than 16 this function returns middle points.
func SizeSixteenOver(number float64) (int, int) {
	w, h := Size()
	if w < 16 || h < 16 {
		return MidPoints()
	}

	if number > 15 || number < 1 {
		return MidPoints()
	}

	w = int(math.Round((float64(w) / 16.0) * number))
	h = int(math.Round((float64(h) / 16.0) * number))

	return w, h
}

// SizeFourOver returns returns Number/4 of the
// width and height of the console based on the number you enter.
// If width of height of console lower than 4 this function returns middle points.
func SizeEightOver(number float64) (int, int) {
	w, h := Size()
	if w < 8 || h < 8 {
		return MidPoints()
	}

	if number > 7 || number < 1 {
		return MidPoints()
	}

	w = int(math.Round((float64(w) / 8.0) * number))
	h = int(math.Round((float64(h) / 8.0) * number))

	return w, h
}

// SizeFourOver returns returns Number/4 of the
// width and height of the console based on the number you enter.
// If width of height of console lower than 4 this function returns middle points.
func SizeFourOver(number float64) (int, int) {
	w, h := Size()
	if w < 4 || h < 4 {
		return MidPoints()
	}

	if number > 3 || number < 1 {
		return MidPoints()
	}

	w = int(math.Round((float64(w) / 4.0) * number))
	h = int(math.Round((float64(h) / 4.0) * number))

	return w, h
}

// Clear clears all font to from console. If it fails it returns error.
func Clear() error {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return errClear
	}
	return nil
}

// Flush refreshes the console to show font changes. If it fails it returns error.
func Flush() error {
	err := termbox.Flush()
	if err != nil {
		return errFlush
	}
	return nil
}

// Init starts termbox to use console. If it fails it returns error.
func Init() error {
	err := termbox.Init()
	if err != nil {
		return errInit
	}
	return nil
}

// Print prints characters to the console
func Print(str string, fgcolor, bgcolor termbox.Attribute, x, y int) {
	for _, r := range str {
		termbox.SetCell(x, y, r, fgcolor, bgcolor)
		x++
	}
	Flush()
}

// Close closes the termbox
func Close() {
	termbox.Close()
}
