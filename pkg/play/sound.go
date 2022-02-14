package play

import (
	"errors"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/mustafanafizdurukan/pomodoro/pkg/logs"
)

var soundNames []string

var (
	errDirNotExist  = errors.New("directory does not exist")
	errDirNotOpened = errors.New("directory could not opened")
)

func sound() (string, error) {
	if len(soundNames) > 0 {
		randNum := rand.Intn(len(soundNames) - 1)
		return soundNames[randNum], nil
	}

	p, err := os.Executable()
	if err != nil {
		logs.ERROR.Println(err)
		return "", err
	}

	basePath := filepath.Join(filepath.Dir(p), "sounds")

	fi, err := os.Stat(basePath)
	if err != nil {
		return "", errDirNotOpened
	}

	if !fi.Mode().IsDir() {
		return "", errDirNotExist
	}

	err = filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// We do not want to break walkdir for specific errors.
			logs.ERROR.Printf("File walk error on path: %s err: %v", path, err)
			return nil
		}

		if d == nil || !d.Type().IsRegular() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".mp3" {
			return nil
		}

		soundNames = append(soundNames, path)

		return nil
	})
	if err != nil {
		return "", err
	}

	randNum := rand.Intn(len(soundNames) - 1)
	return soundNames[randNum], nil
}

func subSound(sound string) {
	for i, s := range soundNames {
		if s == sound {
			soundNames = append(soundNames[:i], soundNames[i+1:]...)
		}
	}
}
