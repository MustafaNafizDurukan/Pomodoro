// Package console implement functions to action about console
package console

import (
	"errors"

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
func MidPoint() (int, int) {
	w, h := Size()
	return w / 2, h / 2
}

// SizeSixteenOver returns returns Number/16 of the
// width and height of the console based on the number you enter.
// If width of height of console lower than 16 this function does not work.
func SizeSixteenOver(number int) (int, int) {
	w, h := Size()
	if w < 16 || h < 16 {
		return 0, 0
	}

	if number > 3 || number < 1 {
		return Size()
	}

	return (w / 16) * number, (h / 16) * number
}

// SizeFourOver returns returns Number/4 of the
// width and height of the console based on the number you enter.
// If width of height of console lower than 4 this function does not work.
func SizeFourOver(number int) (int, int) {
	w, h := Size()
	if w < 4 || h < 4 {
		return 0, 0
	}

	if number > 3 || number < 1 {
		return Size()
	}

	return (w / 4) * number, (h / 4) * number
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
}

// Close closes the termbox
func Close() {
	termbox.Close()
}
