package play

import (
	"errors"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/mustafanafizdurukan/pomodoro/assets"
)

var (
	errOpen   = errors.New("given sound could not be found")
	errDecode = errors.New("given sound could not be decode")
	errPlay   = errors.New("given sound could not be played")
)

// Ring plays random sound that is in sounds directory.
func Ring() error {
	m, err := sound()
	if err != nil {
		return err
	}

	return play(m)
}

func Sound() error {
	m, err := sound()
	if err != nil {
		return err
	}

	return play(m)
}

func play(m string) error {
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

// Warning plays warning messages when any session stopped for a long time.
func Warning(warning assets.Warning) {
	music, err := assets.LoadEmbeddedWarning(warning)
	if err != nil {
		return
	}

	d, err := mp3.NewDecoder(music)
	if err != nil {
		return
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return
	}
	defer c.Close()

	player := c.NewPlayer()
	defer player.Close()

	if _, err := io.Copy(player, d); err != nil {
		return
	}
}
