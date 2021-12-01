// Package font implements functions to get of font informations and print them.
package font

import (
	"time"
	"unicode/utf8"

	"github.com/mustafanafizdurukan/pomodoro/pkg/console"
	"github.com/nsf/termbox-go"
)

type Font struct {
	Text     string
	Fgcolor  termbox.Attribute
	Bgcolor  termbox.Attribute
	Position Position
}

type Position struct {
	X int
	Y int
}

// New returns Font struct
func New(fgcolor termbox.Attribute, bgcolor termbox.Attribute, position *Position) *Font {
	return &Font{
		Fgcolor:  fgcolor,
		Bgcolor:  bgcolor,
		Position: *position,
	}
}

type symbol []string

// Width returns width of given symbol
func (s symbol) Width() int {
	return utf8.RuneCountInString(s[0])
}

// Height returns height of given symbol
func (s symbol) Height() int {
	return len(s)
}

type font []symbol

// Width returns width of given font
func (t font) Width() int {
	w := 0
	for _, s := range t {
		w += utf8.RuneCountInString(s[0])
	}
	return w
}

// Height returns height of given font
func (t font) Height() int {
	return len(t[0])
}

var isCalculationExecuted bool

// Echo prints text as font to the console.
func (f *Font) Echo() {
	font := toFont(f.Text)

	if !isCalculationExecuted {
		f.calculatePoints(font)
		isCalculationExecuted = true
	}

	f.echo(font)
	console.Flush()
}

func toFont(str string) font {
	symbolSlice := make(font, 0)
	for _, r := range str {
		if s, ok := symbols[r]; ok {
			symbolSlice = append(symbolSlice, s)
		}
	}
	return symbolSlice
}

func (f *Font) echo(font font) {
	xSymbol := f.Position.X
	xline, yLine := f.Position.X, f.Position.Y
	for _, s := range font {
		for _, line := range s {
			for _, r := range line {
				termbox.SetCell(xline, yLine, r, f.Fgcolor, f.Bgcolor)
				xline++
			}
			xline = xSymbol
			yLine++
		}
		yLine = f.Position.Y
		xSymbol += s.Width()
		xline = xSymbol
	}
}

func (f *Font) calculatePoints(fo font) {
	x, y := console.MidPoint()
	f.Position.X = x - fo.Width()/2
	f.Position.Y = y - fo.Height()/2
}

var zero = "00:00"

// EchoZero prints 00:00 for 2 seconds
func (f *Font) EchoZero() {
	console.Clear()
	font := toFont(zero)

	f.calculatePoints(font)

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second / 3)
		f.echo(font)
		console.Flush()
		time.Sleep(time.Second / 2)
		console.Clear()
		console.Flush()
	}
}
