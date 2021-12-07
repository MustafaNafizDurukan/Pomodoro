package convert

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringToDate(t *testing.T) {
	var tests = []struct {
		given    string
		expected time.Duration
	}{
		{"3s", 3 * time.Second},
		{"4h4m", (4*time.Hour + 4*time.Minute)},
		{"4m25s", (4*time.Minute + 25*time.Second)},
	}

	for _, test := range tests {
		d, err := StringToDate(test.given)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, d)
	}
}
