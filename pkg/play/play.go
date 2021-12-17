package play

import (
	"errors"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/mustafanafizdurukan/pomodoro/pkg/list"
)

var (
	errOpen   = errors.New("play: given music could not be found")
	errDecode = errors.New("play: given music could not be decode")
	errPlay   = errors.New("play: given music could not be played")
)

func Sound(music string) error {
	if music == "" {
		list.Sound()
	}

	f, err := os.Open(list.Sound())
	if err != nil {
		return errOpen
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return errDecode
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return errPlay
	}

	return nil
}
