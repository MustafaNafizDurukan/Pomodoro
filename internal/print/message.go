package print

import "math/rand"

var messages = []string{
	"Don't look at me",
	"Leave me alone!",
	"Don't give up!",
	"You are almost there!",
	"Go back to your work!",
	"Hang in there!",
	"Don't be distracted!",
	"You can do it!",
	"Stay focused!",
	"Stop surfing the web!",
	"What you plant now, you will harvest later.",
	"Keep up the good work!",
	"Don’t let fear of failure put you off!",
	"Don’t quit yet!",
	"Keep going",
}

// Message returns random message from messages list
func message() string {
	randNum := rand.Intn(len(messages) - 1)
	return messages[randNum]
}
