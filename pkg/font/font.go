// Package font implements functions to get of font informations and print them.
package font

import (
	"unicode/utf8"

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
func (s symbol) width() int {
	return utf8.RuneCountInString(s[0])
}

// Height returns height of given symbol
func (s symbol) height() int {
	return len(s)
}

type font []symbol

// Width returns width of given font
func (f font) width() int {
	w := 0
	for _, s := range f {
		w += utf8.RuneCountInString(s[0])
	}
	return w
}

// Height returns height of given font
func (f font) height() int {
	return len(f[0])
}

// Size returns width and height of given font
func Size(f font) (int, int) {
	return f.width(), f.height()
}

// Echo prints text as font to the console.
func (f *Font) Echo() {
	font := ToFont(f.Text)

	f.echo(font)
}

// ToFont converts given string to font.
func ToFont(str string) font {
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
		xSymbol += s.width()
		xline = xSymbol
	}
}
