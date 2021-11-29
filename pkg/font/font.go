// Package font implements functions to get of font informations and print them.
package font

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

type Font struct {
	text     string
	fgcolor  termbox.Attribute
	bgcolor  termbox.Attribute
	position Position
}

type Position struct {
	X int
	Y int
}

// New returns Font struct
func New(text string, fgcolor termbox.Attribute, bgcolor termbox.Attribute, position *Position) *Font {
	return &Font{
		text:     text,
		fgcolor:  fgcolor,
		bgcolor:  bgcolor,
		position: *position,
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

// Echo prints text as font to the console.
func (f *Font) Echo() {
	font := toFont(f.text)

	f.echo(font)
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
	xSymbol := f.position.X
	xline, yLine := f.position.X, f.position.Y
	for _, s := range font {
		for _, line := range s {
			for _, r := range line {
				termbox.SetCell(xline, yLine, r, f.fgcolor, f.bgcolor)
				xline++
			}
			xline = xSymbol
			yLine++
		}
		yLine = f.position.Y
		xSymbol += s.Width()
		xline = xSymbol
	}
}
