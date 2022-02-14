package play

import (
	"errors"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

var (
	errOpen   = errors.New("given sound could not be found")
	errDecode = errors.New("given sound could not be decode")
	errPlay   = errors.New("given sound could not be played")
)

func Sound() error {
	m, err := sound()
	if err != nil {
		return err
	}

	f, err := os.Open(m)
	if err != nil {
		subSound(m)
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
