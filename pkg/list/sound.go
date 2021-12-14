package list

import (
	"fmt"
	"math/rand"
)

var sounds = []string{
	"discord.mp3",
	"ring.mp3",
	"iphone.mp3",
}

// Sound returns random sound from sounds list
func Sound() string {
	randNum := rand.Intn(len(sounds) - 1)
	return fmt.Sprintf("../sounds/%s", sounds[randNum])
}
